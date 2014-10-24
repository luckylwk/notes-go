package main

import (
	"fmt"
)

type Car struct {
	wheelCount int
	name string
}

func (car Car) numberOfWheels() int {
	return car.wheelCount
} 

type Ferrari struct {
	Car //anonymous field Car
}

type AstonMartin struct {
	Car
}
func (f Ferrari) sayHiToSchumacher() {
	fmt.Println("Hi Schumacher")
}

func (a AstonMartin) sayHiToBond() {
	fmt.Println("Hi Bond, James Bond")
}
func main() {
	f := Ferrari{Car{4, "ferrari"}}
	fmt.Println("A ferrari has this many wheels: ", f.numberOfWheels())
	f.sayHiToSchumacher()

	a := AstonMartin{Car{4, "aston!"}}
	fmt.Println("An Aston Martin has this many wheels: ", a.numberOfWheels())
	a.sayHiToBond()
}