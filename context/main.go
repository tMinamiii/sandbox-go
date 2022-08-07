package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	cancel()

	select {
	case <-ctx.Done():
		fmt.Println("context cancelled")
		os.Exit(1)
	}
	fmt.Println("context has continues")
}
