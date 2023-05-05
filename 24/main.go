package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Line struct {
	Point
}

func (l *Line) GetDistance(firstPoint, secondPoint Point) float64 {
	res := math.Sqrt((math.Pow(firstPoint.x-secondPoint.x, 2)) + math.Pow(firstPoint.y-secondPoint.y, 2))
	return math.Round(res*100) / 100
}

func main() {
	FirstPoint := NewPoint(4.0, 10)
	SecondPoint := NewPoint(10.0, 0)

	var line Line

	fmt.Println(line.GetDistance(FirstPoint, SecondPoint))
}
