package main

import "fmt"

func main() {

	numbers := [...]int{1,2,3,4,5}
	for i, number := range numbers{
		fmt.Printf("%d) %d\n", i, number)
	}

	for _, number := range numbers{
		fmt.Println(number)
	}
}
