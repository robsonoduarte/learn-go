package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name string
	surname string
}

type product struct {
	name string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - $ %.2f", p.name, p.price)
}

func print(p printable) {
	fmt.Print(p.toString())
}

func main() {
	var p printable = person{"Robson", "Duarte"}
	print(p)
	p = product{"Computer", 2.0}
	print(p)
}

