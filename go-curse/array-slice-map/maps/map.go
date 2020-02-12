package main

import "fmt"

func main() {
	var approved = make(map[int]string) // using make to create map
	// set key/value
	approved[1]= "Robson"
	approved[2]= "Ana"
	approved[3]= "Aparecida"
	approved[4]= "Antonio"

	// for range
	for id, name := range approved{
		fmt.Printf("%s (ID: %d)\n", name, id)
	}

	// access value by key
	fmt.Println(approved[1])
	delete(approved, 1) // remove by key
	fmt.Println(approved[1])


	// create/initialize map using inference
	 fees := map[string]float64{
		"Robson": 22.0,
		"Cida": 33.0,
		"Antonio": 44.0,
	}

    fees["Ana"] = 44.0 // adding new key/value

	for name, fee := range fees{
		fmt.Printf("%s (Fee: %f)\n", name, fee)
	}

	// map with map as value
	letters := map[string]map[string]float64{
		"A":{
			"Antonio": 2.2,
			"Aparecida": 3.3,
		},
		"R":{
			"Robson": 33,
			"Ricardo": 33.4,
		},
	}

	for l, m := range letters{
		fmt.Println(l, m)
	}
}
