package main

import (
	"fmt"
	"log"
)

func Input(message, defaultValue string, a ...any) string {
	var in string
	fmt.Printf(message, a...)
	if _, err := fmt.Scanln(&in); err != nil {
		log.Println("Error reading user input: ", err.Error())
		return defaultValue
	}
	if in == "" {
		return defaultValue
	}
	return in
}
