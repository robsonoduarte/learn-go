package main

import "fmt"

func main() {
	x := 20
	fmt.Println(x)

	f := closure()
	f()
}

func closure() func(){
	x := 10
	return func() {
		fmt.Println(x)
	}
}
