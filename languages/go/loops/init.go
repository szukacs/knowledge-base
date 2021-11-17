package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("Outer loop %d inner loop %d\n", i, j)
		}
	}

	i := 0
	for i <10 {
		fmt.Printf("while %d\n",i)
		i++
	}
}
