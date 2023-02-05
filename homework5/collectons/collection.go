package collectons

import (
	"fmt"
	"go-homeworks/homework5/collectons/figures"
	"time"
)

type Collection struct {
	Figures []figures.Figure
	History []string
}

func NewCollection(figures []figures.Figure, history []string) Collection {
	return Collection{
		Figures: figures,
		History: history,
	}
}

func (c *Collection) GetCountFigure() int {
	c.History = append(c.History, fmt.Sprintf("GetCountFigure was called. Time: %s", time.Now()))
	return len(c.Figures)
}

func (c *Collection) AddFigure(figure figures.Figure) {
	c.Figures = append(c.Figures, figure)
	c.History = append(c.History, fmt.Sprintf("AddFigure was called. Time: %s. Arguments: %v", time.Now(), figure))
}

func (c *Collection) GetTotalSquare() float64 {
	var totalSquare float64 = 0
	for _, figure := range c.Figures {
		totalSquare = totalSquare + figure.Square()
	}
	c.History = append(c.History, fmt.Sprintf("GetTotalSquare was called. Time: %s", time.Now()))
	return totalSquare
}

func (c *Collection) GetTotalPerimeter() float64 {
	var totalPerimeter float64 = 0
	for _, figure := range c.Figures {
		totalPerimeter = totalPerimeter + figure.Perimeter()
	}
	c.History = append(c.History, fmt.Sprintf("GetTotalPerimeter was called. Time: %s", time.Now()))
	return totalPerimeter
}

func (c Collection) GetHistoryCall() []string {
	return c.History
}
