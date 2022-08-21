package goslice

import (
	"fmt"
	"testing"
)

func Test_Unshift(t *testing.T) {
	a := []int{1, 2, 3}
	b := Unshift(a, 1)
	fmt.Println(b)
}
