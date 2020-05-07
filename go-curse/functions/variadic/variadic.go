package main

import "fmt"

func main() {
	fmt.Printf("Average => %.2f", average(2.4, 4.6, 7.22))
}

func average(numbers... float64) float64{
	total := 0.0
	for _, number := range numbers{
		total = total + number
	}
	return total / float64(len(numbers))
}
