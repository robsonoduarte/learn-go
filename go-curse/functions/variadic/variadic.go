package main

import "fmt"

// https://gobyexample.com/variadic-functions
func main() {
	fmt.Printf("Average => %.2f", average(2.4, 4.6, 7.22))
	// variadic function with slice
	approved := []string{"Robson", "Ana Mara", "Antonio", "Aparecida"}
	printApproved(approved...)
}

func average(numbers... float64) float64{
	total := 0.0
	for _, number := range numbers{
		total = total + number
	}
	return total / float64(len(numbers))
}

func printApproved(approved ...string) {
	fmt.Println("List of Approved")
	for i, okay := range approved {
		fmt.Printf("%d) %s\n", i+1, okay)
	}
}
