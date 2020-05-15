package main

func main() {
	n := 1
	inc1(n)
	println(n)

	inc2(&n)
	println(n)
}

func inc1(n int){
	n++
}

func inc2(n *int)  {
	*n++
}
