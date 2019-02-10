package main

import "fmt"

func main() {
	a := 10

	if a < 0 {
		fmt.Println("Your value is negative!")
	} else if a < 10 {
		fmt.Println("Your value is a single digit!")
	} else {
		fmt.Println("Your value has multiple digits!")
	}
}
