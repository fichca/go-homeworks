package figures

import "math"

type Square struct {
	Side float64
}

func NewSquare(side float64) Square {
	return Square{Side: side}
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s Square) Square() float64 {
	return math.Pow(s.Side, 2)
}
