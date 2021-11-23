package main

import (
	"fmt"
)

type A struct{}

type As []A

func LoadPtr() *As {
	var a As
	return &a
}

func Load() As {
	var a As
	return a
}

func main() {
	aPtr := LoadPtr()
	fmt.Printf("aPtr: %v\n", aPtr)                         // &[]
	fmt.Printf("*aPtr: %v\n", *aPtr)                       // []
	fmt.Printf("(aPtr == nil): %v\n", aPtr == nil)         // false
	fmt.Printf("(len(*aPtr) == 0): %v\n", len(*aPtr) == 0) // true

	a := Load()
	fmt.Println(a)           // []
	fmt.Println(a == nil)    // true
	fmt.Println(len(a) == 0) // true

	var n []int
	fmt.Println(n)           // []
	fmt.Println(n == nil)    // true
	fmt.Println(len(n) == 0) // true
}
