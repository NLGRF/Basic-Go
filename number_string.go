package main

import "fmt"

func main() {

// Number

fNumber := 60
fSecond := 5

// 65
fmt.Println(fNumber + fSecond)
// 55
fmt.Println(fNumber - fSecond)
// 300
fmt.Println(fNumber * fSecond)
// 12
fmt.Println(fNumber / fSecond)
// 0
fmt.Println(fNumber % fSecond)

// String

p1 := "Mustache"
p2 := "As"

p3 := p1 + " " + p2

// Mustache As
fmt.Println(p3)
// Mus
fmt.Println(p3[0:3])
// 77
fmt.Println(p3[0])
// Mustache As
fmt.Println(p3[0:])

}
