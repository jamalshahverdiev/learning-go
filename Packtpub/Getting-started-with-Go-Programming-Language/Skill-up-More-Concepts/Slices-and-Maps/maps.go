package main

import "fmt"

func main() {
	// m := map[string]string{}
	// m["first"] = "John"
	// m["last"] = "Doe"

	// m := map[string]string{
	// 	"first": "John",
	// 	"last": "Doe",
	// }
	// fmt.Println(m)
	m := make(map[string]string)
	m["first"] = "John"
	m["last"] = "Doe"
	fmt.Println(m["first"])
}
