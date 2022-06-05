package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title  string
	Author struct {
		FirstName string
		LastName  string
	}
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title: "Real World HTTP",
		Author: struct {
			FirstName string
			LastName  string
		}{
			FirstName: "渋川",
			LastName:  "よしき",
		},
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4831111",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book)
}
