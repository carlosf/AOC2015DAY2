package main

import (
	"fmt"
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
	return 2*b.length*b.width + 2*b.width*b.height + 2*b.height*b.length
}

func main() {

}
