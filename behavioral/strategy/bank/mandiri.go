package bank

// Mandiri defines bank Mandiri
type Mandiri struct {
	name             string
	interestRate     float64
	installmentMonth int
	discountPromo    float64
}

// NewBankMandiri initializes new bank Mandiri instance
func NewBankMandiri() *Mandiri {
	return &Mandiri{
		name:             "Mandiri",
		interestRate:     0.025,
		discountPromo:    50000,
		installmentMonth: 12,
	}
}

// GetInterestRate will get bank Mandiri interest rate
func (m *Mandiri) GetInterestRate() float64 {
	return m.interestRate * 100
}

// SetInstallmentMonth will set bank Mandiri installment month period
func (m *Mandiri) SetInstallmentMonth(period int) {
	m.installmentMonth = period
}

// GetBankName returns bank name
func (m *Mandiri) GetBankName() string {
	return m.name
}

// GetInstallmentPeriod returns bank installment period
func (m *Mandiri) GetInstallmentPeriod() int {
	return m.installmentMonth
}

// GetMonthlyInstallment will get bank Mandiri monthly instllment for 12 months
func (m *Mandiri) GetMonthlyInstallment(price float64) float64 {
	return ((price * (1 + m.interestRate)) - m.discountPromo) / float64(m.installmentMonth)
}
