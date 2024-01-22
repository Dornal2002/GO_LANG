package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Println("Enter radius")
	fmt.Scan(&r)

	area := math.Pi * math.Pow(r, 2)
	ans := fmt.Sprintf("%.2f", area)

	fmt.Println("Area of Circle", ans)
}
