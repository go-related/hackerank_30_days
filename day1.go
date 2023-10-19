package main

import (
	"fmt"
)

func (e *Exercises) day1() {
	var word string

	if e.scanner.Scan() {
		word = e.scanner.Text()
	}

	fmt.Println("Hello, World.")
	fmt.Println(word)
}
