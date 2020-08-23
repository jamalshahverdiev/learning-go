package main

import (
	"calculator"
	"fmt"
)

func main() {
	fnum, snum := calculator.MultipleInts()
	fmt.Print(`
1) add
2) divide
3) multiply
4) subtract
Input number of the action to use entered numbers and press 'Enter' button: `)
	var i int
	fmt.Scan(&i)
	if i == 1 {
		fmt.Println("Selected action: 'Add'")
		fmt.Println("The result of 'Addme' method:", calculator.Addme(fnum, snum))
	} else if i == 2 {
		fmt.Println("Selected action: 'Divide'")
		fmt.Println("The result of 'Divide' method:", calculator.Divide(fnum, snum))
	} else if i == 3 {
		fmt.Println("Selected action: 'Multiply'")
		fmt.Println("The result of 'Multiply' method:", calculator.Multiply(fnum, snum))
	} else if i == 4 {
		fmt.Println("Selected action: 'Subtract'")
		fmt.Println("The result of 'Subtract' method:", calculator.Subtract(fnum, snum))
	} else {
		fmt.Println("You can input only digits between 1-4 to use Calculator")
	}
}
