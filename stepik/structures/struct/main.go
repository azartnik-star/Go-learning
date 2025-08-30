package main

import (
	f "fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

/*
	type MultiShape struct {
		shapes []Shape
	}

	func (m *MultiShape) area() float64 {
		var area float64
		for _, s := range m.shapes {
			area += s.area()
		}
		return area
	}
*/
type Person struct {
	Name string
}

func (p *Person) Talk() {
	f.Println("Hi, my name is", p.Name)
}

type Andorid struct {
	Person
	Model string
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 5}
	f.Println(r.area())
	f.Println(c.area())
	a := new(Andorid)
	a.Talk()
	f.Println(totalArea(&c, &r))
	f.Println(r.perimeter())
	f.Println(c.perimeter())
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

/*type Engine struct {
	Noise string
}

func (c *Engine) Start () {
	f.Println(c.Noise, "Engine starting...")
}

type Car struct {
	Engine
}

func main() {
	c := Car{Engine: Engine{Noise:"Vroom"}}
	c.Start()
}*/
