package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName string
	Age int
}
func main() {
	p1 := person {
		FirstName: "Bob",
		LastName: "Tom",
		Age: 33,
	}
	fmt.Println(p1)

	p2 := person {
		FirstName: "Tim",
		LastName: "Tom",
		Age: 22,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println(string(bs))

	
}