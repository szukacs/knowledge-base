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

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello! My name is", p.firstName, p.lastName, "My hobbyis basketball.")
}

func (prog programer) speak() {
	fmt.Println("Hello! My name is", prog.firstName, prog.lastName, "I am a ", prog.language, "developer.")
}

func blaaa(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am a human",h.(person).firstName)
	case programer:
		fmt.Println("I am a human", h.(programer).language)
	}
	//fmt.Println("A am a human", h)
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

 per := person {
	 firstName: "Tom",
	 lastName: "Bob",
 }
 per.speak()
 blaaa(st1)
 blaaa(per)

}