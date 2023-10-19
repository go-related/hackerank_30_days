package day5

import (
	"fmt"
	"github.com/juli/hackerank/30_days/scanner"
	"strconv"
)

type Challange struct {
	Scanner scanner.InputScanner
}

func (e *Challange) Run() {
	defer e.Scanner.Close()
	var T, age int

	T = e.GetInt()

	for i := 0; i < T; i++ {
		age = e.GetInt()
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}

func (e *Challange) GetInt() int {
	result, err := strconv.Atoi(e.Scanner.ReadLine())
	checkError(err)
	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
