package main

import (
	"fmt"
	"sync"

	"github.com/cikupin/design-patterns/creational/singleton/dummy"
)

var (
	dummyInstance dummy.IDummy
	once          sync.Once
)

func initDummy(desc string) dummy.IDummy {
	once.Do(func() {
		dummyInstance = new(dummy.Dummy)
		dummyInstance.SetDescription(desc)
	})
	return dummyInstance
}

func main() {
	initDummy("first dummy")
	fmt.Printf("first initializzation : %s\n", dummyInstance.GetDescription())

	fmt.Println("modify dummy...")
	initDummy("second dummy")
	fmt.Printf("second initializzation : %s\n", dummyInstance.GetDescription())

}
