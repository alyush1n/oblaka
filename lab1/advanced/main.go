package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := InitDB()
	if err != nil {
		log.Fatalf("InitDB error: %s", err.Error())
	}

	fmt.Println("App started!")
	var message string
	if len(os.Args) == 1 {
		message = "Default message"
	} else {
		arguments := os.Args[1:]
		message = arguments[0]
	}

	fmt.Printf("Received message: %s\n", message)

	fmt.Println("Inserting...")
	id, err := InsertMessage(message)
	if err != nil {
		log.Fatalf("Insert message error: %s\n", err.Error())
	}
	fmt.Printf("Done! Inserted message ID: %d\n", id)
}
