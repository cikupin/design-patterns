package main

import (
	"fmt"

	"github.com/cikupin/design-patterns/structural/decorator/addons"
	"github.com/cikupin/design-patterns/structural/decorator/beverages"
)

func main() {
	// House Blend
	houseBlend := beverages.NewHouseBlend("House Blend")
	fmt.Printf("%s : %.1f\n", houseBlend.GetDescription(), houseBlend.Cost())

	houseBlendMilkAddon := addons.NewMilk(houseBlend, "Milk")
	fmt.Printf("%s : %.1f\n", houseBlendMilkAddon.GetDescription(), houseBlendMilkAddon.Cost())

	houseBlendSugarAddon := addons.NewSugar(houseBlendMilkAddon, "Sugar")
	fmt.Printf("%s : %.1f\n", houseBlendSugarAddon.GetDescription(), houseBlendSugarAddon.Cost())

	// Dark Roast
	darkRoast := beverages.NewDarkRoast("Dark Roast")
	fmt.Printf("%s : %.1f\n", darkRoast.GetDescription(), darkRoast.Cost())

	darkRoastSugarAddon := addons.NewSugar(darkRoast, "Sugar")
	fmt.Printf("%s : %.1f\n", darkRoastSugarAddon.GetDescription(), darkRoastSugarAddon.Cost())
}
