package figures

type Parallelogram struct {
	SideA   float64
	SideB   float64
	HeightA float64
}

func NewParallelogram(sideA float64, sideB float64, heightA float64) Parallelogram {
	return Parallelogram{
		SideA:   sideA,
		SideB:   sideB,
		HeightA: heightA,
	}
}
func (p Parallelogram) Perimeter() float64 {
	return 2 * (p.SideA + p.SideB)
}

func (p Parallelogram) Square() float64 {
	return p.HeightA * p.SideA
}
