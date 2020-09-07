package main

import (
	"fmt"
	"os"
)

func main() {
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	presentBoard(board)
}

func presentBoard(b [9]int) {
	for i := range b {
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
			if i == len(b)-1 {
				os.Exit(0)
			}
			fmt.Println("---+---+---")
		} else {
			fmt.Printf("   |")
		}
	}
}
