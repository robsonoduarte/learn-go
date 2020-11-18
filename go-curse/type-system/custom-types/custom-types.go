package main

import "fmt"

// declare the type
type score float64

// and use the receivers to add methods on custom type
func (s score) inside(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}

func applyScore(s score) string {
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
	fmt.Println(applyScore(9))
	fmt.Println(applyScore(6.9))
	fmt.Println(applyScore(2.1))
}

// more about receivers => https://tour.golang.org/methods/4