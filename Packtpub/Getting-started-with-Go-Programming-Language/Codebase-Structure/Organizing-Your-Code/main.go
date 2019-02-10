package main

import (
	"fmt"

	"github.com/mspaulding66/helloworld"
)

func main() {
	fmt.Println(helloworld.GetHelloWorld())
	fmt.Printf("User ID: %d\n", helloworld.GetUserID())
}
