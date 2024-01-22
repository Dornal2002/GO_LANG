package main

import (
	"fmt"
)

func main() {

	var p float32
	fmt.Print("Enter principle ")
	fmt.Scan(&p)

	var rate float32
	fmt.Print("Enter Rate ")
	fmt.Scan(&rate)

	var time float32
	fmt.Print("Enter time period ")
	fmt.Scan(&time)

	si := (p * rate * time) / 100
	result := fmt.Sprintf("%.2f", si)
	fmt.Println("Simple Intrest : ", result)

}
