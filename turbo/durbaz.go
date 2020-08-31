package main

import "fmt"

type Car struct {
	brand   string
	model   string
	bantype string
	color   string
	mileage int
	price   int
	state   State
}

type State int8

const (
	Available State = 1
	Sold      State = 2
)

func main() {
	cardetails := carSoldOrNod()
	fmt.Println(cardetails)
	changeCarsState(cardetails)
	fmt.Println(cardetails)
}

func carSoldOrNod() *Car {
	car := &Car{brand: "Volvo", model: "CX60", bantype: "OffRoad", color: "Black", mileage: 75000, price: 85000, state: State(Available)}
	return car
}

func getCars() []Car {
   
}

func changeCarsState(v *Car) {
	fmt.Printf("v=%p | &v=%p\n", v, &v)
	v.state = 2
}
