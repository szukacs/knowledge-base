package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 5
	fmt.Println(x)
	fmt.Println(len(x))

	y := []int{4, 5, 6, 7, 10, 86}
	fmt.Println(y)

	for i, v := range y {
		fmt.Println(i,v)
	}

	fmt.Println(y)
	fmt.Println(y[3:])
	fmt.Println(y[1:3])

	y = append(y, 88, 112, 126, 523)
	fmt.Println(y)

	z := []int{435, 5353, 6446}

	y = append(y, z...)

	fmt.Println(y)

	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	q := make([]int, 10, 100)
	fmt.Println(q)
	fmt.Println(len(q))
	fmt.Println(cap(q))

	first := []string{"szia1", "szia2", "szia3"}
	second := []string{"hello1", "hello2", "hello3"}

	asd := [] [] string{first, second}
	fmt.Println(asd)
}
