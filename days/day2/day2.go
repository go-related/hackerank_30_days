package day2

import (
	"fmt"
	"github.com/juli/hackerank/30_days/scanner"
	"strconv"
)

type Challange struct {
	Scanner scanner.InputScanner
}

func (e *Challange) Run() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	if e.Scanner.Scan() {
		number, _ := strconv.ParseUint(e.Scanner.Text(), 10, 64)
		i = number + i
	}
	if e.Scanner.Scan() {
		decimalNumber, _ := strconv.ParseFloat(e.Scanner.Text(), 64)
		d += decimalNumber
	}
	if e.Scanner.Scan() {
		s += e.Scanner.Text()
	}

	fmt.Println(fmt.Sprintf("%d", i))
	fmt.Println(fmt.Sprintf("%.1f", d))
	fmt.Println(s)
}
