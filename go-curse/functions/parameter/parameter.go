package main

import "fmt"

func main() {
	result := exec(multiply,3,4)
	fmt.Println(result)
}

func multiply(a, b int) int {
	return a * b
}

func exec(f func(int, int) int, p1, p2 int)  int{
	return f(p1, p2)
}