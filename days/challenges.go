package days

import (
	"github.com/juli/hackerank/30_days/days/day1"
	"github.com/juli/hackerank/30_days/days/day2"
	"github.com/juli/hackerank/30_days/days/day3"
	"github.com/juli/hackerank/30_days/scanner"
)

type DayCallange interface {
	Run()
}

func NewDayOne() DayCallange {
	file := "sources/1.txt"
	return &day1.Challange{
		Scanner: scanner.NewFileScanner(file),
	}
}

func NewDayTwo() DayCallange {
	file := "sources/2.txt"
	return &day2.Challange{
		Scanner: scanner.NewFileScanner(file),
	}
}

func NewDayThree() DayCallange {
	file := "sources/3.txt"
	return &day3.Challange{
		Scanner: scanner.NewFileScanner(file),
	}
}
