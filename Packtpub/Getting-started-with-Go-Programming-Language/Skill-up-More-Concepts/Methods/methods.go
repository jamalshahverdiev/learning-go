package main

import (
	"fmt"
	"time"
)

// func getHello(name string) string {
// 	return fmt.Sprintf("Hello, %s\n", name)
// }
type Person struct {
	First string
	Last  string
	Dob   time.Time
}

func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365)
}

func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.First)
}

func NewPerson(first, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last:  last,
		Dob:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

func main() {
	// if runtime.GOOS == "Linux" {
	// fmt.Println(runtime.GOOS)
	// cmd := exec.Command("ls", "-lah")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// }
	// fmt.Println(getHello(os.Args[1]))
	// p := &Person{
	// 	First: "John",
	// 	Last:  "Doe",
	// }
	p := NewPerson("Jane", "Doe", 1980, 1, 1)
	// p.SayHello()
	fmt.Println(p.GetAge())
}
