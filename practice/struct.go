package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10

	p2 := person{name: "Scan", age: 50}
	p3 := new(person)
	p3.name = "Lee"
	p3.age = 123

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(1)
}
