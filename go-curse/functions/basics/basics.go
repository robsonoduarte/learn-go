package main

import "fmt"

func main() {
	f1()
	f2("Robson", "Ana")

	r1, r2 := f3(), f4("Robson", "Ana")

	fmt.Println(r1)
	fmt.Println(r2)

	r1, r2 = f5()

	fmt.Println(r1, r2)
}

// basic declaration ( without parameters/returns)
func f1() {
	fmt.Println("F1")
}

// with parameters
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

// with return
func f3() string {
	return "F3"
}

// syntax sugar when parameters are of same type
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s, %s", p1, p2)
}

// return multiple values
func f5() (string, string) {
	return "Robson", "Ana"
}

