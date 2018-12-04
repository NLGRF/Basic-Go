package main

import (
	"fmt"
)

func main() {
	result := addition2(5, 5)
	fmt.Println(result * 10)
}

func dosome() {
	fmt.Println("OK")
}

func dosomething(str string) {
	fmt.Println(str)
}

func addition(a int, b int) {
	fmt.Println(a + b)
}

func addition2(a int, b int) int {
	output := (a + b)
	return output
}
