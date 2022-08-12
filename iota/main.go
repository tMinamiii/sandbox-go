package main

import "fmt"

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

func main() {
	// var o CarOption

	// o = SunRoof | HeatedSeat

	// if o&SunRoof != 0 {
	// 	fmt.Println("サンルーフ付き")
	// }
	fmt.Println(GPS)
	fmt.Println(AWD)
	fmt.Println(SunRoof)
	fmt.Println(HeatedSeat)
	fmt.Println(DriverAssist)
}
