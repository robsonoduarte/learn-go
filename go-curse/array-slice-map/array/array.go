package main

import "fmt"

func main() {
	var grades [3]float64 // initialize
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	// traditional loop on Array
	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	average := total / float64(len(grades))

	fmt.Printf("Average is %.2f/n", average)
}
