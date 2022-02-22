/*
 * @Date: 2022-02-22 17:11:27
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-22 17:11:28
 * @FilePath: /moduletest/example/hello.go
 */
package main

import (
	"fmt"
	"log"

	"github.com/NorwayLobster/moduletest/v3/greetings"
	"github.com/NorwayLobster/moduletest/v3/hello"
)

func main() {
	fmt.Println("hello world")
	hello.Hello()

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	log.SetFlags(1)
	log.SetFlags(2)
	log.SetFlags(3)

	// Request a greeting message.
	// message, err := greetings.Hello("")
	message, err := greetings.Hello("xiaowang")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(messages)
}
