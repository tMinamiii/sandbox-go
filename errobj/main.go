package main

import (
	"errors"
	"fmt"
)

// ErrOrig represents task colors
type ErrOrig struct {
	err error
}

func (e *ErrOrig) Error() string {
	return fmt.Sprintf("Original")
}

type ErrSub struct {
	err error
}

func (e *ErrSub) Error() string {
	return fmt.Sprintf("Sub")
}

func main() {
	org := &ErrOrig{
		err: errors.New("abc"),
	}
	sub := &ErrSub{
		err: org,
	}
	detectErr(sub)
}

// detectErr represents
func detectErr(err error) {
	if _, ok := err.(*ErrOrig); ok {
		fmt.Println("Orig is sub")
	}

	if _, ok := err.(*ErrSub); ok {
		fmt.Println("Sub is Sub")
	}
}
