package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter 3 Values Seperated by Comma : ")

	var input string
	fmt.Scanln(&input)

	values := strings.Split(input, ",")

	var p, r, t float64
	fmt.Sscan(values[0], &p)
	fmt.Sscan(values[1], &r)
	fmt.Sscan(values[2], &t)

	si := (p * r * t) / 100
	result := fmt.Sprintf("%.2f", si)
	fmt.Println("Simple Intrest : ", result)
}
