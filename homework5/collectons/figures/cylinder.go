package figures

import "math"

type Cylinder struct {
	Height float64
	Circle
}

func NewCylinder(height float64, circle Circle) Cylinder {
	return Cylinder{
		Height: height,
		Circle: circle,
	}
}

func (c Cylinder) Perimeter() float64 {
	return (c.Radius*2 + c.Height) * 2
}

func (c Cylinder) Square() float64 {
	return math.Pi * c.Radius * (c.Height + c.Radius)
}

func (c Cylinder) Volume() float64 {
	return c.Square() * c.Height
}
