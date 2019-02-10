package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 4}
	// s = append(s, 5)
	// s := make([]int, 10)
	// s[0] = 10

	// This is array example with the predefined size of the array bytes (append method will not work here)
	s := [4]int{1, 2, 3, 4}
	fmt.Println(s)
}
