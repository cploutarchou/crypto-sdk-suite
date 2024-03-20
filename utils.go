package main

import (
	"encoding/json"
	"log"
)

func printJSON(v interface{}) {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error marshaling to JSON:", err)
		return
	}
	log.Println(string(d))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleErrorWithPrint(data interface{}, err error) {
	if err != nil {
		handleError(err)
		return
	}
	printJSON(data)
}
