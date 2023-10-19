package scanner

import (
	"bufio"
	"os"
)

type RealScanner struct {
	scanner *bufio.Scanner
}

func NewStdInScanner() *RealScanner {
	return &RealScanner{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *RealScanner) Scan() bool {
	return r.scanner.Scan()
}

func (r *RealScanner) Text() string {
	return r.scanner.Text()
}

func (r *RealScanner) Err() error {
	return r.scanner.Err()
}
