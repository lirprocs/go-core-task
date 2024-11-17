package main

import (
	"testing"
)

func TestProcessNumbers(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)
	expected := []float64{1, 8, 27, 64, 125, 216, 343, 512, 729, 1000}

	go processNumbers(input, output)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	var result []float64
	for val := range output {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Fatalf("unexpected length: expected %d, result %d", len(expected), len(result))
	}

	for k, v := range expected {
		if v != result[k] {
			t.Errorf("processNumbers(): expected %.2f, result %.2f at key %d", v, result[k], k)
		}
	}
}
