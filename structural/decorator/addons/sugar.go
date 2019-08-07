package addons

import (
	"fmt"

	"github.com/cikupin/design-patterns/structural/decorator/beverages"
)

// Sugar defines struct for sugar addon
type Sugar struct {
	beverages   beverages.IBeverage
	description string
}

// NewSugar will initialize sugar instance
func NewSugar(bev beverages.IBeverage, desc string) *Sugar {
	return &Sugar{
		beverages:   bev,
		description: desc,
	}
}

// SetBeverage will set beverage + sugar
func (s *Sugar) SetBeverage(bev beverages.IBeverage) {
	s.beverages = bev
}

// SetDescription will set sugar's description
func (s *Sugar) SetDescription(desc string) {
	s.description = desc
}

// GetDescription will get sugar's description
func (s *Sugar) GetDescription() string {
	return fmt.Sprintf("%s with Mocha", s.beverages.GetDescription())
}

// Cost will get beverage's cost + sugar
func (s *Sugar) Cost() float64 {
	return (s.beverages.Cost() + 50.0)
}
