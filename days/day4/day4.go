package day4

import (
	"fmt"
	"github.com/juli/hackerank/30_days/scanner"
	"strconv"
	"strings"
)

type Challange struct {
	Scanner scanner.InputScanner
}

func (e *Challange) Run() {
	defer e.Scanner.Close()
	NTemp, err := strconv.ParseInt(strings.TrimSpace(e.Scanner.ReadLine()), 10, 64)
	checkError(err)
	N := int32(NTemp)
	result := "Weird"
	if N%2 == 0 {
		if N >= 2 && N < 5 {
			result = "Not Weird"
		} else if N > 20 {
			result = "Not Weird"
		}
	}
	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
