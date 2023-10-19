package main

import "github.com/juli/hackerank/30_days/scanner"

type Exercises struct {
	scanner scanner.InputScanner
}

func NewExercise_1() *Exercises {
	file := "sources/1.txt"
	return &Exercises{
		scanner: scanner.NewFileScanner(file),
	}
}

func NewExercise_2() *Exercises {
	file := "sources/2.txt"
	return &Exercises{
		scanner: scanner.NewFileScanner(file),
	}
}
