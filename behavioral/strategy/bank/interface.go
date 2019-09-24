package bank

// IBank defines bank interface
type IBank interface {
	GetInterestRate() float64
	SetInstallmentMonth(period int)
	GetBankName() string
	GetInstallmentPeriod() int
	GetMonthlyInstallment(price float64) float64
}
