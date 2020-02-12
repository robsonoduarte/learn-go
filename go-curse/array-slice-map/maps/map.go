package main

import "fmt"

func main() {
	var approved = make(map[int]string)
	approved[1]= "Robson"
	approved[2]= "Ana"
	approved[3]= "Aparecida"
	approved[4]= "Antonio"
	for id, name := range approved{
		fmt.Printf("%s (ID: %d)\n", name, id)
	}

	fmt.Println(approved[1])
	delete(approved, 1)
	fmt.Println(approved[1])
}
