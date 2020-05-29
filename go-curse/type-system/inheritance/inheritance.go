package main

import "fmt"

// Anonymous fields in structs - like object composition  => http://golangtutorials.blogspot.com/2011/06/anonymous-fields-in-structs-like-object.html

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car     // anonymous field
	turboOn bool
}

func main() {
	f := ferrari{
		car{
			"Ferrari 488",
			150,
		},
		false,
	}

	fmt.Println(f.name)
	fmt.Println(f.currentSpeed)
	fmt.Println(f.turboOn)
	fmt.Println(f)
}
