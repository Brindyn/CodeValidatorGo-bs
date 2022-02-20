package main

import (
	"fmt"
	"github.com/Brindyn/CodeValidatorGo-bs/tree/master/standards"
)

func main() {
	// Stores the first input word
	var op string
	// Stores the second input word
	var com string
	// Concatenates the input to make code more readable
	var userInput string
	condition := false
	// Friendly user interface
	fmt.Println()
	fmt.Println("Coding Standards Validator by Brindyn Schultz.")
	fmt.Println("Please type 'val help' for a list of commands.")
	fmt.Println()
	// runs a forever loop
	for !condition {
		// Stores input as pointers for op and com
		fmt.Scan(&op)
		fmt.Scan(&com)
		userInput = (op + " " + com)
		// checks if the first word is val, otherwise outputs an error and retries
		if op == "val" {
			if com != "exit" {
				// passes com to an encapsulated module where it will be executed
				err := standards.FindCommand(com)
				if err != "nil" {
					// gives an error if command is not known
					fmt.Println("Unknown command '" + userInput + "'")
					fmt.Println()
				}
			} else {
				// breaks the forever loop when "val exit" is typed
				fmt.Println("Program quit successfully.")
				fmt.Println()
				break
			}
		} else {
			// gives an error if "val" is not input
			fmt.Println("Unknown command '" + userInput + "'")
			fmt.Println()
		}
	}
}
