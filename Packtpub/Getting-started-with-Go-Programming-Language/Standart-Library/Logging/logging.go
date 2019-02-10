package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var name string

func main() {
	// log.Println("Hello, World!")
	// log.Panic("oops!")
	// log.Fatalln("oops!")
	prefix := fmt.Sprintf("%s: ", os.Args[0])
	// infoLog := log.New(os.Stdout, prefix, log.LstdFlags)
	info, err := os.Create("info.log")
	if err != nil {
		log.Fatalln("Failed to create log file!")
	}
	infoLog := log.New(info, prefix, log.LstdFlags)
	errorLog := log.New(os.Stderr, prefix, log.LstdFlags)
	if name == "" {
		errorLog.Println("No name supplied!")
		flag.Usage()
		os.Exit(1)
		// log.Fatalln("No name supplied!")
		// errorLog.Fatalln("No name supplied!")
	}
	// log.Println("Program Started!")
	// log.Printf("Hello, %s\n", name)
	// log.Println("Program finished!")
	infoLog.Println("Program Started!")
	infoLog.Printf("Hello, %s\n", name)
	infoLog.Println("Program finished!")
}

func init() {
	flag.StringVar(&name, "name", "", "The name to say hello to")
	flag.Parse()
}
