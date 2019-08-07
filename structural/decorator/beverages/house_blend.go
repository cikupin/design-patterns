package beverages

// HouseBlend defines struct for house blend
type HouseBlend struct {
	description string
}

// NewHouseBlend will initialize house blend instance
func NewHouseBlend(desc string) *HouseBlend {
	return &HouseBlend{
		description: desc,
	}
}

// SetDescription will set house blend beverage description
func (hb *HouseBlend) SetDescription(desc string) {
	hb.description = desc
}

// GetDescription will get house blend description
func (hb *HouseBlend) GetDescription() string {
	return hb.description
}

// Cost will get house blend cost
func (hb *HouseBlend) Cost() float64 {
	return float64(250)
}
