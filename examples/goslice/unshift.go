package goslice

func Unshift(s []int, elem int) []int {
	return append([]int{elem}, s...)
}
