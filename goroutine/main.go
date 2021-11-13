package main

import (
	"fmt"
	"sync"
)

const concurrency = 5

var worker *ConcWorker

type ConcWorker struct {
	ch chan func()
	wg *sync.WaitGroup
}

func Start() *ConcWorker {
	var once sync.Once
	once.Do(
		func() {
			var wg sync.WaitGroup
			w := &ConcWorker{
				ch: make(chan func()),
				wg: &wg,
			}
			for i := 0; i < concurrency; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for fn := range w.ch {
						fn()
					}
				}()
			}
			worker = w
		})
	return worker
}

func (c *ConcWorker) Add(job func()) {
	c.ch <- job
}

func (c *ConcWorker) Terminate() {
	close(c.ch)
	c.wg.Wait()
}

type AppendJob struct {
	mu     *sync.Mutex
	result []string
}

func (a *AppendJob) Do() {
	a.mu.Lock()
	a.result = append(a.result, "1")
	a.mu.Unlock()

}

func main() {
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
