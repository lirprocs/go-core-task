package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 2, 4, 6, 8, 10}

	result := sliceExample(slice)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("sliceExample() = %v; expected %v", result, expected)
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}

	result := addElements(slice, 0)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("addElements() = %v; expected %v", result, expected)
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := copySlice(slice)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("copySlice() = %v; expected %v", result, expected)
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result, err := removeElement(slice, len(slice)-1)
	if err != nil {
		t.Errorf("removeElement() returned an error: %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("removeElement() = %v; expected %v", result, expected)
	}

	_, err = removeElement(slice, -1)
	if err == nil {
		t.Error("removeElement() should return an error")
	}

	_, err = removeElement(slice, len(slice))
	if err == nil {
		t.Error("removeElement() should return an error")
	}
}
