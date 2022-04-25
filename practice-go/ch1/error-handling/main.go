package main

import (
	"bufio"
	"fmt"
	"os"
)

func goodHandling() {
	f, err := os.Open("important.txt")
	if err != nil {
		// error handling
	}

	r := bufio.NewReader(f)
	_, err = r.ReadString('\n')
	if err != nil {
		// error handling
	}
	// no error
}

func badHandling() {
	// DONT NEST
	f, err := os.Open("important.txt")
	if err == nil {
		r := bufio.NewReader(f)
		l, err := r.ReadString('\n')
		if err == nil {
			fmt.Println(l)
		}
	}
}
