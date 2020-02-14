package main

import "fmt"

//https://tour.golang.org/basics/7

func main() {
	x, y :=  change(1,2)
	fmt.Println(x,y)
}

func change(p1, p2 int) (second, first int) {
	first = p1
	second = p2
	return
}
