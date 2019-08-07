package beverages

// IBeverage defines interface for beverage
type IBeverage interface {
	SetDescription(desc string)
	GetDescription() string
	Cost() float64
}
