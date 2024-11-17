package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for _, n := range expected[:3] {
			c1 <- n
		}
		close(c1)
	}()

	go func() {
		for _, n := range expected[3:6] {
			c2 <- n
		}
		close(c2)
	}()

	go func() {
		for _, n := range expected[6:] {
			c3 <- n
		}
		close(c3)
	}()

	result := merge(c1, c2, c3)

	var resultSlice []int
	myMap := make(map[int]bool, 9)
	for v := range result {
		resultSlice = append(resultSlice, v)
		myMap[v] = true
	}

	if len(resultSlice) != len(expected) {
		t.Errorf("unexpected length: expected %d, result %d", len(expected), len(resultSlice))
	}

	for _, v := range expected {
		if !myMap[v] {
			t.Errorf("missing value: %d", v)
		}
	}

}
