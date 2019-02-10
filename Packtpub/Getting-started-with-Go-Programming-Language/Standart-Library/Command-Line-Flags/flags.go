package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var debug bool
var name string
var wait time.Duration

func main() {
	//name := flag.String("name", "", "The name to say hello to")

	if name == "" {
		fmt.Println("must add name to use this tool!")
		flag.Usage()
		os.Exit(1)
	}

	if debug {
		fmt.Printf("Going to wait for %v\n", wait)
	}

	time.Sleep(wait)
	fmt.Printf("Hello, %s\n", name)
}

func init() {
	flag.BoolVar(&debug, "debug", false, "Turn on debugging output")
	flag.StringVar(&name, "name", "", "The name to say hello to")
	defaultWait, err := time.ParseDuration("5s")
	if err != nil {
		panic("could not parse default wait time")
	}
	flag.DurationVar(&wait, "wait-time", defaultWait, "Time to wait before print")
	flag.Parse()
}
