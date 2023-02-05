package figures

type Rectangle struct {
	SideA float64
	SideB float64
}

func NewRectangle(sideA float64, sideB float64) Rectangle {
	return Rectangle{
		SideA: sideA,
		SideB: sideB,
	}
}

func (r Rectangle) Perimeter() float64 {
	return (r.SideA + r.SideB) * 2
}

func (r Rectangle) Square() float64 {
	return r.SideA * r.SideB
}
