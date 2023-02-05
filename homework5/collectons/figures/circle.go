package figures

import "math"

type Circle struct {
	Radius float64
}

func NewCircle(Radius float64) Circle {
	return Circle{
		Radius: Radius,
	}
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Square() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
