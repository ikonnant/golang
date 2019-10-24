package main

import (
	"fmt"
	"math"
)

type Coords struct {
	x, y float64
}

type Circle struct {
	center Coords
	radius float64
}

type Triangle struct {
	A, B, C Coords
}

type Shape interface {
	area() float64
}

func main() {
	var m1 map[int]bool
	m2 := map[int]bool{}
	m2[1] = true

	test1 := m2[1]
	test2, ok2 := m2[2]
	_, ok3 := m2[3]

	fmt.Println(m1, m2, test1, test2, ok2, ok3)

	var c1 Coords
	c2 := new(Coords)
	c3 := Coords{x: 11, y: 23}

	fmt.Println("x:", c1.x, ", y:", c1.y)
	fmt.Println("x:", c2.x, ", y:", c2.y)
	fmt.Println("x:", c3.x, ", y:", c3.y)

	fmt.Println(sum(c3))

	circle := Circle{
		center: Coords{0, 10},
		radius: 100,
	}

	triangle := Triangle{
		A: Coords{0, 0},
		B: Coords{10, 0},
		C: Coords{20, 20},
	}

	fmt.Println(circle.area())
	fmt.Println(triangle.area())

	fmt.Println(totalArea(&circle))
}

func sum(c Coords) float64 {
	return c.x + c.y
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t *Triangle) area() float64 {
	return 0.5 * math.Abs((t.A.x-t.B.x)*(t.B.y-t.C.y)-(t.C.x-t.A.x)*(t.A.y-t.B.y))
}

func totalArea(shape Shape) float64 {
	return shape.area()
}
