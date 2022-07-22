package main //must be titled main to work for the go run .

import (
	"fmt"
	"log"
	"example.com/greetings" //module name from the go.mod file in that folder...
)

func logMessage(message string) {
	fmt.Println(message)
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Jeremy")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
			log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	
	// Request a greeting message.
	message, err = greetings.Goodbye("Jeremy")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
			log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}