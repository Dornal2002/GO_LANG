package main

import (
	"errors"
	"fmt"
)

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (rect rectangle) Area() int {
	return rect.length * rect.breadth
}

func (sqr square) Area() int {
	return sqr.side * sqr.side
}

func (rectPerim rectangle) Perimeter() int {
	return 2 * (rectPerim.length + rectPerim.breadth)
}

func (sqPerimeter square) Perimeter() int {
	return 4 * sqPerimeter.side
}

func Print(choice int) (result string, err error) {
	switch choice {
	case 1:
		var length, breadth int
		fmt.Println("Enter Length and Breadth for rectangle : ")
		_, err := fmt.Scanf("%d%d", &length, &breadth)
		if err != nil {
			return "", errors.New("Error!! While reading input for rectangle")
		}
		input := rectangle{
			length:  length,
			breadth: breadth,
		}
		result = fmt.Sprintf("Area : %d\n Perimeter : %d", input.Area(), input.Perimeter())

	case 2:
		var side int
		fmt.Println("Enter Value of Side : ")
		_, err := fmt.Scan(&side)
		if err != nil {
			return "", errors.New("Error !!! While Reading input for square")
		}
		input := square{
			side: side,
		}
		result = fmt.Sprintf("Area : %d\n Perimeter : %d", input.Area(), input.Perimeter())

	default:
		return "", errors.New("*** Plz, Enter Valid Choice...Try Again !! ***")
	}
	return result, nil
}

func main() {
	var choice int

	fmt.Print("\n1.Rectangle\n2.Square\nEnter Your Choice : ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Enter Valid Choice !!", err)
		return
	}

	result, err := Print(choice)
	if err != nil {
		fmt.Println("Error !!! :", err)
		return
	} else {
		fmt.Println(result)
	}
}
