package main

import (
	"fmt"
)

func main() {
	fmt.Println("Input Your Number:")
	var input float64
	fmt.Scanf("%f", &input)
	condition := input > 2
	if condition {
		fmt.Println("Worked")
	} else {
		fmt.Println("Not Worked")
	}
}
