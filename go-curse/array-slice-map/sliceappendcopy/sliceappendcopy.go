package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}

	var s1 []int

	//a1 = append(a1, 4, 5, 6) don't compile append need a slice not an array

	s1 = append(s1, 4, 5, 6)
	fmt.Println(a1,s1)

	s2 := make([]int, 2)
	copy(s2,s1)
	fmt.Println(s2)
}
