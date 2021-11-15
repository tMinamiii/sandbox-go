package worker

import (
	"sync"
)

const concurrency = 5

var worker *ConcurrencyWorker
var once sync.Once

type ConcurrencyWorker struct {
	ch chan func()
	wg *sync.WaitGroup
}

func Start() *ConcurrencyWorker {
	once.Do(
		func() {
			var wg sync.WaitGroup
			w := &ConcurrencyWorker{
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

func (c *ConcurrencyWorker) Add(job func()) {
	c.ch <- job
}

func (c *ConcurrencyWorker) Terminate() {
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
