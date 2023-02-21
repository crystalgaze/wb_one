package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func main() {
	a := NewPoint(-10, -10)
	b := NewPoint(10, 10)
	fmt.Println(Distance(a, b))
}
