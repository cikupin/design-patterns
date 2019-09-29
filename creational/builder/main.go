package main

import (
	"github.com/cikupin/design-patterns/creational/builder/car"
)

func main() {
	yarisMT := car.NewYarisBuilder().
		SetModel("yaris MT").
		SetFrontBrake("disc").
		SetFrontDiscBrakeType("solid").
		SetRearBrake("drum").
		SetTransmission("manual").
		Build()
	yarisMT.PrintSpecs()

	yarisAT := car.NewYarisBuilder().
		SetModel("yaris AT").
		SetFrontBrake("disc").
		SetFrontDiscBrakeType("ventilated").
		SetRearBrake("drum").
		SetRearDiscBrakeType("ventilated").
		SetTransmission("automatic").
		SetAutomaticType("cvt").
		Build()
	yarisAT.PrintSpecs()
}
