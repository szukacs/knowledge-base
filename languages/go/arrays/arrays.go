package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 5
	fmt.Println(x)
	fmt.Println(len(x))

	y := []int{4, 5, 6, 7}
	fmt.Println(y)

	for i, v := range y {
		fmt.Println(i,v)
	}
}
