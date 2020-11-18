package main

import "fmt"

type score float64

func (s score) inside(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}

func x(s score) string {
	switch {
	case s.inside(9.0, 10.0):
		return "A"
	case s.inside(7.0, 8.99):
		return "B"
	case s.inside(5.0, 7.99):
		return "C"
	case s.inside(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(x(9))
	fmt.Println(x(6.9))
	fmt.Println(x(2.1))
}
