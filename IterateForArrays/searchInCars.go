package main

import "fmt"

// CarData struct will be used as constructor of our Mock data.
type CarData struct {
	brand   string
	model   string
	bantype string
	color   string
	mileage uint32
	price   uint32
	year    uint16
	state   State
}

// State will be used inside of the struct
type State int8

// Enumerated const will be used to dynamically change state inside of the struct
const (
	Available State = 1
	Sold      State = 2
)

var allCarsData = []CarData{
	{brand: "Volvo", model: "6", bantype: "Sedan", color: "White", mileage: 7000, price: 15000, year: 2016, state: State(Available)},
	{brand: "Mercedes", model: "C200", bantype: "Sedan", color: "Metallic", mileage: 7000, price: 10000, year: 2018, state: State(Available)},
	{brand: "BMW", model: "745i", bantype: "Sedan", color: "Black", mileage: 10000, price: 100000, year: 2014, state: State(Sold)},
	{brand: "Mitsubishi", model: "Lancer", bantype: "Sedan", color: "White", mileage: 1000, price: 15000, year: 2012, state: State(Available)},
	{brand: "Hyundai", model: "Accent", bantype: "Sedan", color: "White", mileage: 123000, price: 12000, year: 2000, state: State(Available)},
	{brand: "Mazda", model: "6", bantype: "Sedan", color: "Red", mileage: 2, price: 40000, year: 2020, state: State(Available)},
	{brand: "Audi", model: "A9", bantype: "Sedan", color: "Black", mileage: 1, price: 75000, year: 2011, state: State(Available)},
	{brand: "Wolkswagen", model: "Jetta", bantype: "Sedan", color: "Bkack", mileage: 106000, price: 16500, year: 2006, state: State(Available)},
	{brand: "Toyota", model: "Rav4", bantype: "OffRoad", color: "White", mileage: 1, price: 48000, year: 2020, state: State(Available)},
	{brand: "Kia", model: "Sorento", bantype: "OffRoad", color: "Black", mileage: 10000, price: 20000, year: 2013, state: State(Available)},
}

func searchCar(carsDataArr []CarData) []string {
	var carBrands []string
	for _, elem := range carsDataArr {
		if elem.year >= 2015 {
			fmt.Printf("Car brand: %s | Car year: %d\n", elem.brand, elem.year)
			carBrands = append(carBrands, elem.brand)
		}
	}
	return carBrands
}

func main() {
	fmt.Println(searchCar(allCarsData))
}

