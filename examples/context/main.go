package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func timeout(ctx1 context.Context, i int) {
	ctx2, cancel := context.WithTimeout(ctx1, 5*time.Second)
	defer cancel()
	fmt.Println("sleep...")
	time.Sleep(6 * time.Second)

	fmt.Printf("i=%d ctx2 %p : ", i, ctx2)
	fmt.Println(ctx2.Deadline())
	select {
	case <-ctx2.Done():
		fmt.Println("child context cancelled")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("parent context cancelled")
			fmt.Println(ctx.Deadline())
			os.Exit(1)
		}
	}()

	timeout(ctx, 1)
	timeout(ctx, 2)
	timeout(ctx, 3)
	timeout(ctx, 4)
	timeout(ctx, 5)
	timeout(ctx, 6)
	timeout(ctx, 7)
	timeout(ctx, 8)
	timeout(ctx, 9)
}
