package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar um var ( declara e atribuindo)
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência = ", area)

	// blocos de const e var

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	fmt.Println(1, b, c, d)

	// na mesma linha com N var
	var e, f bool = true, false

	fmt.Println(e, f)

	// na mesma linha com a forma reduzia
	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
