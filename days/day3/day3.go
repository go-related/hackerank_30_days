package day3

import (
	"fmt"
	"github.com/juli/hackerank/30_days/scanner"
	"math"
	"strconv"
	"strings"
)

type Challange struct {
	Scanner scanner.InputScanner
}

func (e *Challange) Run() {
	defer e.Scanner.Close()
	meal_cost, err := strconv.ParseFloat(strings.TrimSpace(e.readLine()), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(e.readLine()), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(e.readLine()), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)

}
func (e *Challange) readLine() string {
	if e.Scanner.Scan() {
		return e.Scanner.Text()
	}
	return ""
}

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	total_tip := meal_cost * float64(tip_percent) / 100
	total_tax := meal_cost * float64(tax_percent) / 100

	total := math.Round(meal_cost + total_tip + total_tax)
	fmt.Println(fmt.Sprintf("%.0f", total))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
