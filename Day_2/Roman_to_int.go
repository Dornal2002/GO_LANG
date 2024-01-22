package main

import (
	"fmt"
	"strings"
)

func Roman_to_int(inp string) string {

	var val int
	res := 0
	prevVal := -1
	flag := false

	for i := len(inp) - 1; i >= 0; i-- {
		switch inp[i] {
		case 'I':
			val = 1
		case 'V':
			val = 5
		case 'X':
			val = 10
		case 'L':
			val = 50
		case 'C':
			val = 100
		case 'D':
			val = 500
		case 'M':
			val = 1000
		default:
			flag = true
		}

		if prevVal > val {
			res -= val
		} else {
			res += val
		}

		prevVal = val
	}

	if flag {
		return "Invalid Roman Number !!"
	} else {
		output := fmt.Sprintf("Decimal number for given Roman Number is : %d", res)
		return output
	}
}

func main() {

	var input string
	fmt.Println("Enter Roman Value to Find it's Decimal Value : ")
	fmt.Scan(&input)

	input = strings.ToUpper(input)
	result := Roman_to_int(input)
	fmt.Println(result)
}
