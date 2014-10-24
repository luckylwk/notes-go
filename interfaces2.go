package main 
// providing "evolutionary design"

import (
	"fmt"
)

//define the datastructure
type Bus struct {
	l, b, h int
	rows, seatsPerRow int
}

//define a real world abstraction that could use the data

type Cuboider interface {
	CubicVolume() int
}

//Implement methods to work on data
func (bus Bus) CubicVolume() int {
	return bus.l * bus.b * bus.h
}

type PublicTransporter interface {
	PassengerCapacity() int
}

func (bus Bus) PassengerCapacity() int {
	return bus.rows * bus.seatsPerRow
}

/*
A possible scenario:
A law is put in place that requires a certain amount of space for each passenger on a bus. 
Now we have a new situation where the two distinct interfaces set above need to align to meet
this new law
*/
type PersonalSpaceLaw interface {
	IsCompliantWithLaw() bool
}
func (b Bus) IsCompliantWithLaw() bool {
	return b.CubicVolume() / b.PassengerCapacity() >= 3
}

func main() {
	b := Bus {
		l: 10,
		b: 6,
		h: 3,
		rows: 10,
		seatsPerRow: 5 }
	fmt.Println("Cubic Volume of bus: ", b.CubicVolume())
	fmt.Println("Maximum number of passengers: ", b.PassengerCapacity())
	fmt.Println("Is Compliant with Law? ", b.IsCompliantWithLaw())
}