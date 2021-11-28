package main

import "fmt"

func main() {

	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())


	{
		name := "Bob"

		fmt.Println(name)
	}
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}