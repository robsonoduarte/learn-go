package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}

	for {
		fmt.Println("forever...")
		time.Sleep(time.Second * 5)
	}
}
