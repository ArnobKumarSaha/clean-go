package main

import (
	"fmt"
	called "github.com/Arnobkumarsaha/golang/basics/calledpackage"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// This block will call the Hellos function declared in calledPackage/greetings.go
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, err := called.Hellos(names)
	if err != nil {
		log.Fatal(err) // log.Fatal function call os.Exit(1) internally.
	}
	fmt.Println("Hellos : ", message)

	// This block will call the Hello function declared in calledPackage/greeting.go
	anotherMessage, err := called.Hello("Arnob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello : ", anotherMessage)
}

// The last line on go.mod file was created when syncing
// In package name, only characters, digits and underscore are allowed.
