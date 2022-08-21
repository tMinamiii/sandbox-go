package goslice

func Rotate(s []int) []int {
	elem := s[0]
	s = append(s[1:], elem)
	return s
}
