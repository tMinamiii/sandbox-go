package main

import "fmt"

func Unshift(s []int, elem int) []int {
	return append([]int{elem}, s...)
}

func main() {
	a := []int{1, 2, 3}
	b := Unshift(a, 1)
	fmt.Println(b)
}
