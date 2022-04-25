package main

// error interface
type error interface {
	Error() string
}

func New(text string) error {
	return &errorString{text}
}

// Error() stringを実装ているので error インターフェースを満たす
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
