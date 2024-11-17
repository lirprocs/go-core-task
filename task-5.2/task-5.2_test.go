package main

import (
	"testing"
)

func TestRandom(t *testing.T) {
	ch := make(chan int)
	num := 100
	result := make([]int, 0, num)

	go random(ch, num)

	for val := range ch {
		result = append(result, val)
	}

	if len(result) != num {
		t.Errorf("expected %d numbers, got %d", num, len(result))
	}

	for _, v := range result {
		if v >= num || v < 0 {
			t.Errorf("value %d is out of range [0, 100)", v)
		}
	}
}
