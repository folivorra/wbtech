package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками на плоскости.
//Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
//Расстояние рассчитывается по формуле между координатами двух точек.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(2.0, 3.0)

	fmt.Println("distance:", point1.Distance(point2))
}
