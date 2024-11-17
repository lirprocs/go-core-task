package main

import (
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	wg := NewMyWaitGroup()
	delta := 5

	wg.Add(delta)
	if wg.counter != int32(delta) {
		t.Errorf("testAdd(): %d; expected %d", wg.counter, delta)
	}
	wg.Add(delta)
	if wg.counter != int32(delta)*2 {
		t.Errorf("testAdd(): %d; expected %d", wg.counter, delta*2)
	}
}

func TestDone(t *testing.T) {
	wg := NewMyWaitGroup()
	delta := 5

	wg.Add(delta)
	wg.Done()
	if wg.counter != int32(delta)-1 {
		t.Errorf("testDone(): %d; expected %d", wg.counter, delta-1)
	}
}

func TestWait(t *testing.T) {
	var mu sync.Mutex
	wg := NewMyWaitGroup()
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var result []int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			mu.Lock()
			result = append(result, val)
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	myMap := make(map[int]bool, len(expected))
	for _, v := range result {
		myMap[v] = true
	}

	if len(result) != len(expected) {
		t.Errorf("unexpected length: expected %d, result %d", len(expected), len(result))
	}

	for _, v := range expected {
		if !myMap[v] {
			t.Errorf("missing value: %d", v)
		}
	}
}
