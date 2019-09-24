package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/cikupin/design-patterns/behavioral/strategy/bank"
)

// Loan defines loan instance
type Loan struct {
	amount      float64
	institution bank.IBank
}

// SetBank will set loan amount
func (l *Loan) SetBank(newInstitution bank.IBank) {
	l.institution = newInstitution
}

// CalculateInstallment will print installment detail
func (l *Loan) CalculateInstallment() {
	p := message.NewPrinter(language.Indonesian)

	p.Printf("Bank name : %s\n", l.institution.GetBankName())
	p.Printf("Loan amount : Rp %.2f\n", l.amount)
	p.Printf("Interest rate : %.2f%%\n", l.institution.GetInterestRate())
	p.Printf("Installment periond (month) : %d\n", l.institution.GetInstallmentPeriod())
	p.Printf("Monthly installment : Rp %.2f\n", l.institution.GetMonthlyInstallment(l.amount))
	p.Println()
}

func main() {
	loan := Loan{
		amount:      26000000,
		institution: bank.NewBankMandiri(),
	}

	loan.CalculateInstallment()

	// change installment to BCA
	bca := bank.NewBankBCA()
	loan.SetBank(bca)
	loan.CalculateInstallment()

	bca.SetInstallmentMonth(6)
	loan.CalculateInstallment()

	// change installment to BNI
	bni := bank.NewBankBNI()
	bni.SetInstallmentMonth(3)
	loan.SetBank(bni)
	loan.CalculateInstallment()
}
