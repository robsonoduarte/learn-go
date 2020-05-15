package main

func main() {
	println(factorial(5))
}

func factorial(n uint) uint{
	switch  {
	case n == 0:
		return 1
	default:
		return n * factorial(n - 1)
	}
}
