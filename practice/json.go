package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title  string
	Status status
}
type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func ExampleTask_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title:"Buy Milk", "Status":2}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
}

func main() {
	fmt.Println("Hi")
}
