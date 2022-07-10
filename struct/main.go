package main

import (
	"fmt"
)

type A struct {
	name string
}

func (a *A) SetPtr(name string) {
	a.name = name
}

func (a A) Set(name string) {
	a.name = name
}

func main() {
	a := &A{name: "minami"}
	a.SetPtr("hoge")
	fmt.Println(a.name)
	a.Set("piyo")
	fmt.Println(a.name)
}
