package main

import (
	"fmt"
	"strconv"
)

func (e *Exercises) day2() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	if e.scanner.Scan() {
		number, _ := strconv.ParseUint(e.scanner.Text(), 10, 64)
		i = number + i
	}
	if e.scanner.Scan() {
		decimalNumber, _ := strconv.ParseFloat(e.scanner.Text(), 64)
		d += decimalNumber
	}
	if e.scanner.Scan() {
		s += e.scanner.Text()
	}

	fmt.Println(fmt.Sprintf("%d", i))
	fmt.Println(fmt.Sprintf("%.1f", d))
	fmt.Println(s)
}
