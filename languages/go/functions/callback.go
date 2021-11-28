package main

import "fmt"

func main () {
	nums := []int{1,2,3,4,5,6,7,8,9}
	s := sum(nums...)
	fmt.Println(s)

	fmt.Println(even(sum, nums...))
}
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v %2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}