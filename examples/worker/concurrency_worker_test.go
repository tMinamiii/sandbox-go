package worker

import (
	"fmt"
	"sync"
	"testing"
)

func Test_ConcurrencyWorker(t *testing.T) {
	w := Start()
	mu := sync.Mutex{}
	job := &AppendJob{
		mu: &mu,
	}
	var expected []string
	const count = 30

	for i := 1; i < count; i++ {
		expected = append(expected, "1")
	}

	for i := 1; i < count; i++ {
		w.ch <- job.Do
	}

	w.Terminate()
	if len(job.result) != len(expected) {
		fmt.Printf("wrong expected = %d, actual = %d\n", len(expected), len(job.result))
	} else {
		fmt.Println("complete")
	}
}
