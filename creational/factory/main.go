package main

import (
	"log"

	"github.com/cikupin/design-patterns/creational/factory/cars"
)

func main() {
	carFactory := cars.NewCarFactory()

	agya, err := carFactory.GetCar("agya MT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	agya.PrintSpecs()

	xpander, err := carFactory.GetCar("xpander exceed AT")
	if err != nil {
		log.Fatalln(err.Error())
	}
	xpander.PrintSpecs()

	lamborghini, err := carFactory.GetCar("lamborghini aventador")
	if err != nil {
		log.Fatalln(err.Error())
	}
	lamborghini.PrintSpecs()
}
