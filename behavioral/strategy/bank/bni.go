package bank

// BNI defines bank BNI
type BNI struct {
	name             string
	interestRate     float64
	installmentMonth int
}

// NewBankBNI initializes new bank BNI instance
func NewBankBNI() *BNI {
	return &BNI{
		name:             "BNI",
		interestRate:     0.006,
		installmentMonth: 12,
	}
}

// GetInterestRate will get bank BNI interest rate
func (bni *BNI) GetInterestRate() float64 {
	return bni.interestRate * 100
}

// SetInstallmentMonth will set bank BNI installment month period
func (bni *BNI) SetInstallmentMonth(period int) {
	bni.installmentMonth = period
}

// GetBankName returns bank name
func (bni *BNI) GetBankName() string {
	return bni.name
}

// GetInstallmentPeriod returns bank installment period
func (bni *BNI) GetInstallmentPeriod() int {
	return bni.installmentMonth
}

// GetMonthlyInstallment will get bank BNI monthly instllment for 12 months
func (bni *BNI) GetMonthlyInstallment(price float64) float64 {
	return (price * (1 + bni.interestRate)) / float64(bni.installmentMonth)
}
