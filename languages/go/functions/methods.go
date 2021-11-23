package main

import "fmt"


type person struct {
	firstName string
	lastName string
}

type programer struct {
	person
	language string
}

func (prog programer) speak() {
	fmt.Println("Hello! My name is", prog.firstName, prog.lastName)
}

func main() {
 st1 := programer {
	 person: person{
		 firstName: "Geri",
		 lastName: "Smith",
	 },
	 language: "Go",
 }

 fmt.Println(st1)
 st1.speak()
}