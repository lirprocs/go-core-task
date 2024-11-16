package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"banana", "date"}

	result := intersection(slice1, slice2)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("intersection() = %v; expected %v", result, expected)
	}
}
