package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.name, p.age)
}

type ByAge []Person

func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
 numbers := []int {4, 2, 6, 1, 7, 10, 17, 12}
 strings := []string {"Bob", "Tom", "Adam"}

 fmt.Println(numbers)
 fmt.Println(strings)

 sort.Ints(numbers)
 sort.Strings(strings)

 fmt.Println(numbers)
 fmt.Println(strings)

 fmt.Println("-------------------------")

 p1 := Person{"Tom", 22}
 p2 := Person{"Bob", 43}
 p3 := Person{"Adam", 11}
 p4 := Person{"Mike", 74}

 people := []Person{p1, p2, p3, p4}
 
 fmt.Println(people)
 sort.Sort(ByAge(people))
 fmt.Println(people)
	
}