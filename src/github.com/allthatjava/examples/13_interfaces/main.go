package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64{
	return r.width * r.height
}

func getArea(s Shape) float64{
	return s.area()
}

func main(){

	c := Circle{x:0, y:0, radius:5}
	r := Rectangle{width: 10, height:20}

	fmt.Printf("Circle aread:%f\n", getArea(c))
	fmt.Printf("Rectangle aread:%f", getArea(r))


}
