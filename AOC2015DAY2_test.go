package main

import "testing"

func TestBox_getSurfaceArea(t *testing.T) {
	type fields struct {
		length int
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{name: "2x3x4",
			fields: fields{2, 3, 4},
			want:   58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				length: tt.fields.length,
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := b.getSurfaceArea(); got != tt.want {
				t.Errorf("getSurfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
