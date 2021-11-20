package main

import "fmt"

func main() {
 m:=map[string]int{
	 "Geri": 22,
	 "Pityu": 33,
 }

 fmt.Println(m)

 fmt.Println(m["Pityu"])
 fmt.Println(m["Bran"])

 v, ok := m["Bran"]

 fmt.Println(v)
 fmt.Println(ok)

 m["Magdi"] = 33


 if v, ok := m["Geri"]; ok {
	 fmt.Println("found it:)", v)
 }

 for k, v :=range m {
	 fmt.Println( k, v)
 }

 delete(m, "Bandi")
 if v, ok := m["Pityu"]; ok {
	 fmt.Println("value: " , v)
	 delete(m, "Pityu")
 }

 fmt.Println(m)
}
