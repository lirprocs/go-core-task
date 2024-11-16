package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expectedSlice := []int{64, 3}

	resultSlice, resultBool := intersection(a, b)
	if !reflect.DeepEqual(expectedSlice, resultSlice) && resultBool != true {
		t.Errorf("intersection() = %v; expected %v", resultSlice, expectedSlice)
	}

	c := []int{0, 1, 55, 88}
	expectedSlice2 := []int{}
	resultSlice2, resultBool2 := intersection(a, c)
	if !reflect.DeepEqual(expectedSlice2, resultSlice2) && resultBool2 != false {
		t.Errorf("intersection() = %v; expected %v", resultSlice2, expectedSlice2)
	}
}
