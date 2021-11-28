package main

import "fmt"

func main() {
	defer yollolo()
	hello("Greg")
	message := bye("Geri")
	fmt.Println(message)
	value1, value2 := blabla("John", "Smidth")
	fmt.Println(value1)
	fmt.Println(value2)
	numbers := []int{2, 3, 4, 6, 7, 15, 19}
	sum(numbers...)
	sum(3, 4, 6)
	
	func(num int) {
		fmt.Println("Helloooooo", num)
	}(42)

	myFunct := func() {
		fmt.Println("Hello from my fuction")
	}

	myFunct()
	
	secretVar := secret()
	fmt.Printf("%T", secretVar)
	fmt.Println()
	
	fmt.Println(secret()())
}

func secret() func() int {
	return func() int {
		return 300
	}
}

func yollolo() {
	fmt.Println("Yollolo")
}

func hello(name string) {
	fmt.Println("Hello from fuction", name)
}

func bye(name string) string {
	return fmt.Sprint("Bye ", name)
}


func blabla(firstName, lastName string) (string, bool){
	msg := fmt.Sprint(firstName, " ", lastName, `, says "Hello"`)
	saysTruth := false
	return msg, saysTruth
}

func sum(magicNums ...int) int {
	fmt.Println(magicNums)
	fmt.Printf("%T\n", magicNums)
	sum := 0
/*
	for i, v := range magicNums {
		fmt.Print(i)
		fmt.Println()
		sum += v
	}
	*/

	for _, v := range magicNums {
		sum += v
	}
	fmt.Println(len(magicNums))
	fmt.Println(cap(magicNums))
	fmt.Println("Sum: ", sum)
	return sum
}