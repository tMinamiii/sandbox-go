package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("slice: %v\n", slice)
	d := unshift(slice, 4)
	fmt.Printf("unshit result: %v\n", d)
	fmt.Printf("slice: %v", slice)
	e := unshift(d, 5)
	fmt.Printf("unshit result: %v\n", e)
	fmt.Printf("slice: %v\n", slice)
}

func unshift(s []int, elem int) []int {
	return append([]int{elem}, s...)
}
