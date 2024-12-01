package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		expected []int
	}{
		{[]int{2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		c1 := make(chan int)
		c2 := make(chan int)
		c3 := make(chan int)
		go func() {
			for _, n := range tt.expected[:3] {
				c1 <- n
			}
			close(c1)
		}()

		go func() {
			for _, n := range tt.expected[3:6] {
				c2 <- n
			}
			close(c2)
		}()

		go func() {
			for _, n := range tt.expected[6:] {
				c3 <- n
			}
			close(c3)
		}()
		result := merge(c1, c2, c3)
		resultArr := make([]int, len(tt.expected))
		for v := range result {
			resultArr = append(resultArr, v)
		}
		if reflect.DeepEqual(resultArr, tt.expected) {
			t.Errorf("Add() = %d; want %d", resultArr, tt.expected)
		}
	}

}
