package main

import "fmt"

// CarData struct will be used as constructor of our Mock data.
type CarData struct {
	brand   string
	model   string
	banType string
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
	{brand: "Volvo", model: "6", banType: "Sedan", color: "White", mileage: 7000, price: 15000, year: 2016, state: State(Available)},
	{brand: "Mercedes", model: "C200", banType: "Sedan", color: "Metallic", mileage: 7000, price: 10000, year: 2018, state: State(Available)},
	{brand: "BMW", model: "745i", banType: "Sedan", color: "Black", mileage: 10000, price: 100000, year: 2014, state: State(Sold)},
	{brand: "Mitsubishi", model: "Lancer", banType: "Sedan", color: "White", mileage: 1000, price: 15000, year: 2012, state: State(Available)},
	{brand: "Hyundai", model: "Accent", banType: "Sedan", color: "White", mileage: 123000, price: 12000, year: 2000, state: State(Available)},
	{brand: "Mazda", model: "6", banType: "Sedan", color: "Red", mileage: 2, price: 40000, year: 2020, state: State(Available)},
	{brand: "Audi", model: "A9", banType: "Sedan", color: "Black", mileage: 1, price: 75000, year: 2011, state: State(Available)},
	{brand: "Wolkswagen", model: "Jetta", banType: "Sedan", color: "Bkack", mileage: 106000, price: 16500, year: 2006, state: State(Available)},
	{brand: "Toyota", model: "Rav4", banType: "OffRoad", color: "White", mileage: 1, price: 48000, year: 2020, state: State(Available)},
	{brand: "Kia", model: "Sorento", banType: "OffRoad", color: "Black", mileage: 10000, price: 20000, year: 2013, state: State(Available)},
}

func searchCarByYear(start, end uint16) []CarData {
	var cars []CarData
	for _, c := range allCarsData {
		if c.year >= start && c.year <= end {
			cars = append(cars, c)
		}
	}
	return cars
}

func searchCarByColor(color string) []CarData {
	var cars []CarData
	for _, c := range allCarsData {
		if c.color == color {
			cars = append(cars, c)
		}
	}
	return cars
}

func searchCarByBrand(b string) []CarData {
	var cars []CarData
	for _, c := range allCarsData {
		if c.brand == b {
			cars = append(cars, c)
		}
	}
	return cars
}

func main() {
	foundCars := searchCarByYear(2015, 2018)
	for _, c := range foundCars {
		fmt.Println(c)
	}
}
