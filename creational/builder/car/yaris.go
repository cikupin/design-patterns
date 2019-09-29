package car

import "fmt"

// Yaris defines toyota yaris property
type Yaris struct {
	Model              string
	FrontBrake         string
	FrontDiscBrakeType string
	RearBrake          string
	RearDiscBrakeType  string
	Transmission       string
	AutomaticType      string
}

// PrintSpecs will print toyota yaris specs
func (y *Yaris) PrintSpecs() {
	fmt.Printf("Model : %s\n", y.Model)
	fmt.Printf("Front brake : %s\n", y.FrontBrake)
	fmt.Printf("Front disc brake type : %s\n", y.FrontDiscBrakeType)
	fmt.Printf("Rear brake : %s\n", y.RearBrake)

	if y.RearDiscBrakeType != "" {
		fmt.Printf("Rear disc brake type : %s\n", y.RearDiscBrakeType)
	}

	fmt.Printf("Transmission : %s\n", y.Transmission)

	if y.AutomaticType != "" {
		fmt.Printf("Automatic type : %s\n", y.AutomaticType)
	}
	fmt.Println()
}
