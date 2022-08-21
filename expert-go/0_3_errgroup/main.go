package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

// golang.org/x/sync/errgroupパッケージのerrgroup.Group注5 は基本的にsync.WaitGroupと同じで
// すが、並列に実行する関数からエラーを返すことができます。
func main() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		n := 1
		eg.Go(func() error {
			return do(n)
		})
	}

	if err := eg.Wait(); err != nil {
		os.Exit(1)
	}
}

func do(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("err")
	}
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)

	return nil
}
