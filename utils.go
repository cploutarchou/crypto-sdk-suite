// Package main contains utility functions for printing JSON and handling errors.
package main

import (
	"encoding/json"
	"log"
)

// printJSON prints the JSON representation of the given value with indentation.
//
// The function takes an interface{} as input and marshals it to JSON using json.MarshalIndent.
// If there's an error during marshaling, it logs the error and returns.
// Otherwise, it logs the JSON string.
func printJSON(v any) {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error marshaling to JSON:", err)
		return
	}
	log.Println(string(d))
}

// handleError logs the given error using log.Fatal if it's not nil.
//
// The function takes an error as input and checks if it's not nil.
// If the error is not nil, it logs the error using log.Fatal.
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// handleErrorWithPrint handles the given error and prints the data if there's no error.
//
// The function takes an interface{} and an error as input.
// If the error is not nil, it calls handleError to log the error.
// If the error is nil, it calls printJSON to print the data as JSON.
func handleErrorWithPrint(data any, err error) {
	if err != nil {
		handleError(err)
		return
	}
	printJSON(data)
}
