package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func (p *person) setFullName(fullName string) {
	n := strings.Split(fullName, " ")
	p.name = n[0]
	p.lastName = n[1]
}

func main() {
	p := person{
		"Robson",
		"Duarte",
	}

	fmt.Println(p.getFullName())

	p.setFullName("Ana Mara")
	fmt.Println(p.getFullName())
}
