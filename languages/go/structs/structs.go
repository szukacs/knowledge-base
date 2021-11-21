package main

import "fmt"

type person struct {
	name string
	age int
}

type programer struct {
	person
	language string
}

func main() {
	p1 := person{
		name: "Geri",
		age: 22,
	}

	prog1 := programer{
		person: person{
			name: "Steve",
			age: 27,
		},
		language: "go",
	}

	fmt.Println(prog1)
	fmt.Println(prog1.name, prog1.age, prog1.language)
	fmt.Println(p1)
	fmt.Println(p1.name)

	char := struct {
		name string
		elementType string
	}{
		name : "Charmander",
		elementType: "fire",
	}

	fmt.Println(char)
}
