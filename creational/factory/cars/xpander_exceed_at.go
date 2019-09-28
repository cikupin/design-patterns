package cars

import "fmt"

// XpanderExceedAT ...
type XpanderExceedAT struct {
	ABS              bool
	EBD              bool
	FrontBrake       string
	FrontDiscType    string
	RearBrake        string
	StabilityControl bool
	Transmission     string
	AutomaticType    string
}

// NewXpanderExceedAT ...
func NewXpanderExceedAT(opt Options) *XpanderExceedAT {
	return &XpanderExceedAT{
		ABS:              opt.ABS,
		EBD:              opt.EBD,
		FrontBrake:       opt.FrontBrake,
		FrontDiscType:    opt.FrontDiscType,
		RearBrake:        opt.RearBrake,
		StabilityControl: opt.StabilityControl,
		Transmission:     opt.Transmission,
		AutomaticType:    opt.AutomaticType,
	}
}

// PrintSpecs ...
func (a *XpanderExceedAT) PrintSpecs() {
	fmt.Println("Car model : xpander exceed AT")
	fmt.Printf("ABS : %t\n", a.ABS)
	fmt.Printf("EBD : %t\n", a.EBD)
	fmt.Printf("Front brake : %s\n", a.FrontBrake)
	fmt.Printf("Front disc type : %s\n", a.FrontDiscType)
	fmt.Printf("Rear brake : %s\n", a.RearBrake)
	fmt.Printf("Stability control : %t\n", a.StabilityControl)
	fmt.Printf("Transmission : %s\n", a.Transmission)
	fmt.Printf("Automatic type : %s\n\n", a.AutomaticType)
}
