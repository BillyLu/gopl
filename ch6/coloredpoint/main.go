package main

import "image/color"
import "fmt"

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	c := ColoredPoint{Color: red}
	fmt.Printf("c %#v\n", c)
}
