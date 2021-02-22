package main

import (
	"fmt"
	"math"
)

func main() {

	var r geometry = rectangle1{
		width:  30,
		height: 20,
	}
	fmt.Println(r.area1())
	fmt.Println(r.perimeter())

	var c geometry = circle{
		radius: 10,
	}
	fmt.Println(c.area1())
	fmt.Println(c.perimeter())
}

type geometry interface {
	area1() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle1 struct {
	height, width float64
}

func (r rectangle1) area1() float64 {
	return r.width * r.height
}

func (r rectangle1) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area1() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
