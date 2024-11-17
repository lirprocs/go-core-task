package main

import (
	"fmt"
	"sync/atomic"
)

type MyWaitGroup struct {
	counter int32
	done    chan struct{}
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{done: make(chan struct{})}
}

func (wg *MyWaitGroup) Add(delta int) {
	val := atomic.AddInt32(&wg.counter, int32(delta))

	if val < 0 {
		panic("negative counter in MyWaitGroup")
	}

	if val == 0 {
		close(wg.done)
	}
}

func (wg *MyWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *MyWaitGroup) Wait() {
	if atomic.LoadInt32(&wg.counter) == 0 {
		return
	}
	<-wg.done
}

func main() {
	wg := NewMyWaitGroup()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
