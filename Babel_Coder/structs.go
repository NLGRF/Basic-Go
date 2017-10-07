package main

import "fmt"

type human struct {
  name string
  age int
}

/*
somchai := human{name: "Somchai", age: 23}
somsree := human{name: "Somsree", age: 32}
*/

// printInfo จะผูกติดกับ human
// สังเกตมีการระบุ human เข้าไปก่อนหน้าชื่อเมธอด
func (h human) printInfo() {
  fmt.Println(h.name, h.age)
}

func main() {
  somchai := human{name: "Somchai", age: 23}
  somchai.printInfo() // Somchai 23
}
