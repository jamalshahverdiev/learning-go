package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Create FILE and write some data to this file
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	panic("unable to create file")
	// }
	// defer f.Close()
	// cnt, err := f.WriteString("Hello, World!")
	// if err != nil {
	// 	panic("unable to write file")
	// }
	// fmt.Printf("Wrote %d bytes\n", cnt)

	// Read data from file
	// f, err := os.Open("output.txt")
	// if err != nil {
	// 	panic("unable to open file")
	// }
	// defer f.Close()
	// buf := make([]byte, 64)
	// cnt, err := f.Read(buf)
	// if err != nil {
	// 	panic("unable to read file")
	// }
	// fmt.Printf("Read %d bytes\n", cnt)
	// fmt.Println(string(buf[:cnt]))

	// Create and write to file with the IOUtil
	// err := ioutil.WriteFile("out.txt", []byte("Hello, World!"), 0644)
	// if err != nil {
	// 	panic("unable to write file")
	// }
	buf, err := ioutil.ReadFile("out.txt")
	if err != nil {
		panic("unable to read file")
	}
	fmt.Println(string(buf))
}
