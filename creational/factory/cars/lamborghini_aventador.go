package cars

import "fmt"

// LamborghiniAventador ...
type LamborghiniAventador struct {
	ABS              bool
	EBD              bool
	FrontBrake       string
	FrontDiscType    string
	RearBrake        string
	RearDiscType     string
	StabilityControl bool
	Transmission     string
	AutomaticType    string
}

// NewLamborghiniAventador ...
func NewLamborghiniAventador(opt Options) *LamborghiniAventador {
	return &LamborghiniAventador{
		ABS:              opt.ABS,
		EBD:              opt.EBD,
		FrontBrake:       opt.FrontBrake,
		FrontDiscType:    opt.FrontDiscType,
		RearBrake:        opt.RearBrake,
		RearDiscType:     opt.RearDiscType,
		StabilityControl: opt.StabilityControl,
		Transmission:     opt.Transmission,
		AutomaticType:    opt.AutomaticType,
	}
}

// PrintSpecs ...
func (a *LamborghiniAventador) PrintSpecs() {
	fmt.Println("Car model : lamborghini aventador")
	fmt.Printf("ABS : %t\n", a.ABS)
	fmt.Printf("EBD : %t\n", a.EBD)
	fmt.Printf("Front brake : %s\n", a.FrontBrake)
	fmt.Printf("Front disc type : %s\n", a.FrontDiscType)
	fmt.Printf("Rear brake : %s\n", a.RearBrake)
	fmt.Printf("Rear disc brake : %s\n", a.RearDiscType)
	fmt.Printf("Stability control : %t\n", a.StabilityControl)
	fmt.Printf("Transmission : %s\n", a.Transmission)
	fmt.Printf("Automatic type : %s\n\n", a.AutomaticType)
}
