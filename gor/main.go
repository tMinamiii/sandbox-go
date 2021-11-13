package main

import (
	"fmt"
	"sync"
)

const concurrency = 5

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range c {
				fmt.Println(msg)
			}
		}()
	}
	for i := 1; i < 10; i++ {
		c <- fmt.Sprintf("hello%d", i)
	}
	close(c)
	wg.Wait()
	fmt.Println("complete")
}

func worker(c chan string, done chan struct{}) {
	for msg := range c {
		fmt.Println(msg)
	}
	done <- struct{}{}
}
