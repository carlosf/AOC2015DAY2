package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(str string) {
	fmt.Println(str)
}

type Box struct {
	length int
	width  int
	height int
}

func (b Box) getSurfaceArea() int {
	minside := math.Min(float64(b.length*b.width), float64(b.width*b.height))
	minside = math.Min(minside, float64(b.height*b.length))
	fmt.Println("minside is: " + strconv.Itoa(int(minside)))
	return 2*b.length*b.width + 2*b.width*b.height + 2*b.height*b.length + int(minside)
}

func main() {
	p("Start AOC2015DAY2.go")
	//open input.txt file for

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalfeet := 0

	for {
		if !scanner.Scan() {
			break
		}
		p(scanner.Text())
		boxDim := strings.Split(scanner.Text(), "x")
		boxlength, _ := strconv.Atoi(boxDim[0])
		boxwidth, _ := strconv.Atoi(boxDim[1])
		boxheight, _ := strconv.Atoi(boxDim[2])
		p("Box length is: " + strconv.Itoa(boxlength))
		p("Box width is: " + strconv.Itoa(boxwidth))
		p("Box height is: " + strconv.Itoa(boxheight))
		myBox := Box{length: boxlength, width: boxwidth, height: boxheight}
		p("Box surface area is: " + strconv.Itoa(myBox.getSurfaceArea()) + " square feet")
		totalfeet = totalfeet + myBox.getSurfaceArea()

	}
	p("Box size is: " + strconv.Itoa(totalfeet))
	p("End AOC2015DAY2.go")

}
