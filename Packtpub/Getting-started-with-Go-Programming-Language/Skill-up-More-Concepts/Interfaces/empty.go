package main

import "fmt"

func main() {
	var x interface{}
	// x = "Hello, World"
	x = 100
	// s, ok := x.(string)
	// if !ok {
	// 	panic("NO!")
	// }
	// fmt.Println(s)
	switch x.(type) {
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String!")
	}
}
