package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Tick")
		time.Sleep(1 * time.Second)
		//break
	}
}
