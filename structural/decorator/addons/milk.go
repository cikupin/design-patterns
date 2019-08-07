package addons

import (
	"fmt"

	"github.com/cikupin/design-patterns/structural/decorator/beverages"
)

// Milk defines struct for milk addon
type Milk struct {
	beverages   beverages.IBeverage
	description string
}

// NewMilk will initialize milk instance
func NewMilk(bev beverages.IBeverage, desc string) *Milk {
	return &Milk{
		beverages:   bev,
		description: desc,
	}
}

// SetBeverage will set beverage + milk
func (m *Milk) SetBeverage(bev beverages.IBeverage) {
	m.beverages = bev
}

// SetDescription will set milk's description
func (m *Milk) SetDescription(desc string) {
	m.description = desc
}

// GetDescription will get milk's description
func (m *Milk) GetDescription() string {
	return fmt.Sprintf("%s with Milk", m.beverages.GetDescription())
}

// Cost will get beverage's cost + milk
func (m *Milk) Cost() float64 {
	return (m.beverages.Cost() + 100.0)
}
