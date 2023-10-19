package day1

import (
	"fmt"
	"github.com/juli/hackerank/30_days/scanner"
)

type Challange struct {
	Scanner scanner.InputScanner
}

func (e *Challange) Run() {
	defer e.Scanner.Close()
	var word string

	if e.Scanner.Scan() {
		word = e.Scanner.Text()
	}

	fmt.Println("Hello, World.")
	fmt.Println(word)
}
