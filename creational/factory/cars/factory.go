package cars

import (
	"errors"
)

// Options defines car options
type Options struct {
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

// ICar defines car interface
type ICar interface {
	PrintSpecs()
}

// CarFactory defines car
type CarFactory struct{}

// NewCarFactory generate car factory instance
func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

// GetCar gets car instance
func (c *CarFactory) GetCar(model string) (ICar, error) {
	var carSpecs Options
	var car ICar

	switch model {
	case "agya MT":
		carSpecs = Options{
			ABS:              false,
			EBD:              false,
			FrontBrake:       "disc",
			FrontDiscType:    "solid",
			RearBrake:        "drum",
			StabilityControl: false,
			Transmission:     "manual",
		}
		car = NewAgyaMT(carSpecs)

	case "xpander exceed AT":
		carSpecs = Options{
			ABS:              true,
			EBD:              true,
			FrontBrake:       "disc",
			FrontDiscType:    "solid",
			RearBrake:        "drum",
			StabilityControl: false,
			Transmission:     "automatic",
			AutomaticType:    "conventional",
		}
		car = NewXpanderExceedAT(carSpecs)

	case "lamborghini aventador":
		carSpecs = Options{
			ABS:              true,
			EBD:              true,
			FrontBrake:       "disc",
			FrontDiscType:    "ventilated",
			RearBrake:        "disc",
			RearDiscType:     "ventilated",
			StabilityControl: true,
			Transmission:     "automatic",
			AutomaticType:    "dual clutch",
		}
		car = NewLamborghiniAventador(carSpecs)

	default:
		return nil, errors.New("wrong car")
	}

	return car, nil
}
