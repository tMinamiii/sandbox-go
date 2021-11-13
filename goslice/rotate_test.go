package goslice

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Rotate(t *testing.T) {
	s := []int{1, 2, 3}
	actual := Rotate(s)
	want := []int{2, 3, 1}

	if reflect.DeepEqual(actual, want) {
		_ = fmt.Errorf("Wrong! actual=%v, want=%v", actual, want)
	}
}
