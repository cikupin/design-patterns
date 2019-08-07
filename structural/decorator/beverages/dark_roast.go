package beverages

// DarkRoast defines struct for dark roast
type DarkRoast struct {
	description string
}

// NewDarkRoast will initialize dark roast instance
func NewDarkRoast(desc string) *DarkRoast {
	return &DarkRoast{
		description: desc,
	}
}

// SetDescription will set dark roast beverage description
func (dr *DarkRoast) SetDescription(desc string) {
	dr.description = desc
}

// GetDescription will get dark roast description
func (dr *DarkRoast) GetDescription() string {
	return dr.description
}

// Cost will get dark roast cost
func (dr *DarkRoast) Cost() float64 {
	return float64(300)
}
