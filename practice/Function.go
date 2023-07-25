package main

import (
	"fmt"
)

func returnTrue() string {
	return "a" + "b"
}

func main() {
	fmt.Printf("%s", returnTrue())
}
