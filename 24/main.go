package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func FindDistance(p1 *Point, p2 *Point) float64 {
	dx := math.Abs(p1.x - p2.x)
	dy := math.Abs(p1.y - p2.y)
	fmt.Println(dx, dy)

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(3.0, 4.0)

	fmt.Println(FindDistance(p1, p2))
}
