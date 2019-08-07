package addons

import (
	"github.com/cikupin/design-patterns/structural/decorator/beverages"
)

// IAddOn defines interface for addon
type IAddOn interface {
	SetBeverage(bev beverages.IBeverage)
	SetDescription(desc string)
	GetDescription() string
	Cost() float64
}
