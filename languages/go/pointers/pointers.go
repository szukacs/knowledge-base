package main

import "fmt"

func main() {
x := 43

fmt.Println(x)
fmt.Println(&x)

fmt.Printf("%T\n", x)
fmt.Printf("%T\n", &x)

var b *int = &x
fmt.Println(b)
fmt.Println(*b)

*b = 44
fmt.Println(x)
}
