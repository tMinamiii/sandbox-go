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

func LoadEmptyPtr() *As {
	var a *As
	return a
}

func Load() As {
	var a As
	return a
}

func main() {
	aPtr := LoadPtr()
	fmt.Println(aPtr)            // &[]
	fmt.Println(*aPtr)           // []
	fmt.Println(aPtr == nil)     // false
	fmt.Println(len(*aPtr) == 0) // true

	emptyPtr := LoadEmptyPtr()
	fmt.Println(emptyPtr)        //nil
	fmt.Println(emptyPtr == nil) //true
	// fmt.Println(len(*emptyPtr) == 0) //Panic

	a := Load()
	fmt.Println(a)           // []
	fmt.Println(a == nil)    // true
	fmt.Println(len(a) == 0) // true

	var n []int
	fmt.Println(n)           // []
	fmt.Println(n == nil)    // true
	fmt.Println(len(n) == 0) // true
}
