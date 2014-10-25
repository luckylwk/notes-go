package main

import ( 
	"fmt"
)

type Vehicle interface {
	printNumberOfWheels() int
	anotherWayToPrintWheelCount()
}

type Car struct {
	cost int
}

type Ferrari struct {
	Car
	colour string
	numberOfWheels int
}

func (f Ferrari) printNumberOfWheels() int {
	return f.numberOfWheels
}
func (f Ferrari) anotherWayToPrintWheelCount() {
	fmt.Println(f.numberOfWheels)
}

type Ford struct {
	Car
	name string
}

func main() {
	ferrari := Ferrari{Car{200}, "red", 4}
	ford := Ford{Car{100}, "focus"}
	fmt.Println("Ferrari is", ferrari.colour, " and the Ford is a ", ford.name)

	v := Vehicle(ferrari)

	fmt.Println("The nuber of wheels on the vehicle is", v.printNumberOfWheels())
	v.anotherWayToPrintWheelCount()
}