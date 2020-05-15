package main

func main() {
	println(multiply(3,3))
}


func multiply(x uint, y uint) uint{
	defer println("The End")
	println("multiplying...")
	result := x * y
	return result
}
