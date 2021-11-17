package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, go to learn! :D")
	foo()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("Outer loop %d inner loop %d", i, j)
		}
	}
}
var z = "hello"
var t string = "helloYolo"
var g = `hello "Yollo"

and

HI`

type hug int

var b hug
func foo() {
	fmt.Println("yollolo")
	x := 47
	b = 38
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(t)
	fmt.Printf("%T\n",t)
	fmt.Println(g)
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	y = int(b)

	fmt.Println(y)
	fmt.Printf("%T\n",y)


}
