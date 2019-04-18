package main

import "fmt"

// valor default dos tipos
func main() {
	var a int
	var b float32
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
