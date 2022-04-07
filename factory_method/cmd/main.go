package main

import (
	"fmt"
	"patterns/factory_method/knights"
	"patterns/factory_method/pirates"
)

func main() {
	piratesCreator := &pirates.PiratesGang{}
	knightsCreator := &knights.KnightsPlatoon{}

	for _, mem := range append(piratesCreator.CreateSquad(10, 10), knightsCreator.CreateSquad(10, 10)...) {
		fmt.Printf("%v %v\n", mem.Name(), mem.Health())
	}
}
