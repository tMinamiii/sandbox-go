package generics

import "fmt"

type Struct struct {
	t interface{}
}

func (s Struct) String() string {
	return fmt.Sprintf("%v", s.t)
}

func String[T any](s T) string {
	return fmt.Sprintf("%v", s)
}
