package main

import (
	"encoding/json"
	"fmt"
)

	
type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main () {
s := `[{"FirstName":"Bob","LastName":"Tom","Age":33},{"FirstName":"Tim","LastName":"Tom","Age":22}]`
bs := []byte(s)
fmt.Printf("%T\n", s)
fmt.Printf("%T\n", bs)

people := []person{}

err :=json.Unmarshal(bs, &people)
if err != nil {
	fmt.Println(err)
}
fmt.Println("all of the data" , people)

for i, v := range people {
	fmt.Println("PERSON -", i, v)
}
}