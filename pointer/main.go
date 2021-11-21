package main

import (
	"fmt"
)

type A struct {
	name string
}

type As []A

func Load() *As {
	var a *As
	return a
}

func main() {
	a := Load()
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(a == nil)
	fmt.Println(len(*a) == 0)
}
