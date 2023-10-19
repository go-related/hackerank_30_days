package main

import (
	"fmt"
	"log"
)

func Exercise1() {
	var word string

	_, err := fmt.Scan(&word)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Hello, World.")
	fmt.Println(word)
}
