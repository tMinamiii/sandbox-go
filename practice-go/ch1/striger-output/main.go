//go:generate enumer -type=CarOption -json
//go:generate enumer -type=CarType -json

package main

type (
	CarOption uint64
	CarType   int
)

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)
