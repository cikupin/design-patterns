package bank

// BCA defines bank BCA
type BCA struct {
	name             string
	interestRate     float64
	installmentMonth int
	adminFee         float64
}

// NewBankBCA initializes new bank BCA instance
func NewBankBCA() *BCA {
	return &BCA{
		name:             "BCA",
		interestRate:     0.017,
		adminFee:         35000,
		installmentMonth: 12,
	}
}

// GetInterestRate will get bank BCA interest rate
func (b *BCA) GetInterestRate() float64 {
	return b.interestRate * 100
}

// SetInstallmentMonth will set bank BCA installment month period
func (b *BCA) SetInstallmentMonth(period int) {
	b.installmentMonth = period
}

// GetBankName returns bank name
func (b *BCA) GetBankName() string {
	return b.name
}

// GetInstallmentPeriod returns bank installment period
func (b *BCA) GetInstallmentPeriod() int {
	return b.installmentMonth
}

// GetMonthlyInstallment will get bank BCA monthly instllment for 12 months
func (b *BCA) GetMonthlyInstallment(price float64) float64 {
	return ((price * (1 + b.interestRate)) + b.adminFee) / float64(b.installmentMonth)
}
