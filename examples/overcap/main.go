package main

import "fmt"

func main() {
	a := make([]int64, 0, 10)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	fmt.Println(a)
	for _, v := range a {
		fmt.Println(v)
	}

}
