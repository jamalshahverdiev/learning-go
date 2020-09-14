package main

import "fmt"

type computer struct {
	brand      string
	cpu        string
	hdd        uint16
	screensize uint8
	cmemory    uint16
}

type whiteCollection interface {
	monitor(b string, i, s uint8)
	motherboard(b, ch, cp string)
	memory(s, r uint16)
}

func main() {
	size := "SSD"
	switch size {
	case "SSD":
		fmt.Printf("Better disk speed in: %v\n", size)
	case "HDD":
		fmt.Printf("Normal disk speed in: %v\n", size)
	default:
		fmt.Printf("%v | unknown type: %T\n", size, size)
	}

	var w whiteCollection
	w = &computer{brand: "Dell", cpu: "intel_Ci7", hdd: 700, screensize: 15, cmemory: 64}
	w.monitor("HP", 15, 24)
	w.motherboard("Gigabyte", "Intel", "i7")
	w.memory(700, 64)
}

func describe(i *interface{}) {
	fmt.Printf("%v | %T\n", i, i)
}

func add(a, b int) int {
	return a + b
}

func (d *computer) monitor(brand string, inch, size uint8) {
	fmt.Printf("Monitor brand: %s | Inch: %d | Screen size: %d\n", brand, inch, size)
}

func (d *computer) motherboard(brand, chipset, cpu string) {
	fmt.Printf("Motherbard: %s | Chipset: %s | CPU: %s\n", brand, chipset, cpu)
}

func (d *computer) memory(disk, ram uint16) {
	fmt.Printf("HDD size: %d | RAM size: %d\n", disk, ram)
}

func intReturn(sizeofstr string) int {
	return len(sizeofstr)
}

