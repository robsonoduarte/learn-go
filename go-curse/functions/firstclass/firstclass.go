package main

import "fmt"

//https://golang.org/doc/codewalk/functions/

func main() {
	fmt.Println(add(2,3))

	sub := func(a, b int) int{
		return a - b
	}

	fmt.Println(sub(2, 3))
}


var add = func(a, b int)int {
	return  a + b
}
