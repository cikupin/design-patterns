package cars

import (
	"fmt"
)

// AgyaMTSpec ...
type AgyaMTSpec struct {
	ABS              bool
	EBD              bool
	FrontBrake       string
	FrontDiscType    string
	RearBrake        string
	StabilityControl bool
	Transmission     string
}

// NewAgyaMT ...
func NewAgyaMT(opt Options) *AgyaMTSpec {
	return &AgyaMTSpec{
		ABS:              opt.ABS,
		EBD:              opt.EBD,
		FrontBrake:       opt.FrontBrake,
		FrontDiscType:    opt.FrontDiscType,
		RearBrake:        opt.RearBrake,
		StabilityControl: opt.StabilityControl,
		Transmission:     opt.Transmission,
	}
}

// PrintSpecs ...
func (a *AgyaMTSpec) PrintSpecs() {
	fmt.Println("Car model : agya MT")
	fmt.Printf("ABS : %t\n", a.ABS)
	fmt.Printf("EBD : %t\n", a.EBD)
	fmt.Printf("Front brake : %s\n", a.FrontBrake)
	fmt.Printf("Front disc type : %s\n", a.FrontDiscType)
	fmt.Printf("Rear brake : %s\n", a.RearBrake)
	fmt.Printf("Stability control : %t\n", a.StabilityControl)
	fmt.Printf("Transmission : %s\n\n", a.Transmission)
}
