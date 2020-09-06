package main

import (
	"fmt"
	"strings"
)

func main() {
	maximalStars := inputMaximalStarNumber()
	lineNum := 1
	for i := 1; i <= maximalStars; i++ {
		if mod(i, 2) > 0 && i >= 1 {
			spaces := strings.Repeat(" ", maximalStars-lineNum)
			lineNum = lineNum + 1
			stars := strings.Repeat("*", i)
			fmt.Println(spaces, stars)
		}
	}
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func inputMaximalStarNumber() int {
	fmt.Print(`
Welcome to the Yolka!
Please 'Enter' maximal number count of the starts to prepare Yolka.`)
	var f int
	fmt.Print("\nEnter maximal 'Yolka' stars: ")
	fmt.Scan(&f)
	return f
}

