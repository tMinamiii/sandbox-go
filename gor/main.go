package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	done := make(chan struct{})
	// defer close(c)
	go worker(c, done)
	for i := 1; i < 10; i++ {
		c <- fmt.Sprintf("hello%d", i)
	}
	close(c)
	<-done
	fmt.Println("complete")

}
func worker(c chan string, done chan struct{}) {
	for msg := range c {
		fmt.Println(msg)
	}
	done <- struct{}{}
}
