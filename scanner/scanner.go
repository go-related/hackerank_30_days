package scanner

type InputScanner interface {
	Scan() bool
	Text() string
	Err() error
}
