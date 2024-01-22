package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func (area rectangle) Area() int {
	return area.length * area.breadth
}

func (perimeter rectangle) Perimeter() int {
	return 2 * (perimeter.length + perimeter.breadth)
}

func main() {
	var l, b int

	fmt.Println("Enter Length And Width For Rectangle")

	_, err := fmt.Scanf("%d%d", &l, &b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if l <= 0 || b <= 0 {
		fmt.Println("Error: Length and breadth must be positive values")
		return
	}

	val := rectangle{
		length:  l,
		breadth: b,
	}

	fmt.Println("Area Of Rectangle:", val.Area(), "\nPerimeter of Rectangle:", val.Perimeter())
}
