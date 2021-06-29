package main

import (
	"fmt"
	"math"
)

// Написать программу нахождения расстояния между 2 точками, которые представление в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Set(x float64, y float64) {
	p.x, p.y = x, y
}

func (p *Point) Get() (x float64, y float64) {
	return p.x, p.y
}

func (p1 *Point) Range(p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(10, 10)
	fmt.Println(p1.Range(p2))
}
