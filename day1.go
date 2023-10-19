package main

import (
	"bufio"
	"fmt"
	"os"
)

func Exercise1() {
	var word string
	// to read a input with space in go you can't use the fmt directly(  fmt.Scanln(&word))
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		word = scanner.Text()
	}

	fmt.Println("Hello, World.")
	fmt.Println(word)
}
