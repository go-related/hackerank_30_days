package scanner

import (
	"bufio"
	"fmt"
	"os"
)

type FileScanner struct {
	fileName string
	scanner  *bufio.Scanner
}

func NewFileScanner(fileName string) *FileScanner {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	//defer file.Close()
	return &FileScanner{
		fileName: fileName,
		scanner:  bufio.NewScanner(file),
	}
}

func (r *FileScanner) Scan() bool {
	return r.scanner.Scan()
}

func (r *FileScanner) Text() string {
	return r.scanner.Text()
}

func (r *FileScanner) Err() error {
	return r.scanner.Err()
}
