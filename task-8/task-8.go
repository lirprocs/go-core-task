package main

import "fmt"

func processNumbers(input <-chan uint8, output chan<- float64) {
	for num := range input {
		output <- float64(num) * float64(num) * float64(num)
	}
	close(output)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go processNumbers(input, output)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	for val := range output {
		fmt.Printf("%.2f", val)
	}
}
