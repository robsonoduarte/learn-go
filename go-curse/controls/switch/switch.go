package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(testGrade(9.9))
	fmt.Println(testGrade(6.9))
	fmt.Println(testGrade(-1.0))


	fmt.Println(_type("Hello"))
	fmt.Println(_type(func() {}))
	fmt.Println(_type(time.Hour))


	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("AM")
	case t.Hour() > 12:
		fmt.Println("PM")
	}
}

func testGrade(grade float64) string {
	var g = int(grade)
	switch g {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid Grade"
	}
}

func _type(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "func"
	default:
		return "type not found"
	}
}
