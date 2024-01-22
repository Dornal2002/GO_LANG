package main

import "fmt"

type AnyType interface{}

type Hello string

func AcceptAnything(inp AnyType) string {
	var result string
	switch inp.(type) {
	case int:
		result = fmt.Sprintf("This is a Value of type Integer, '%d'", inp)
	case string:
		result = fmt.Sprintf("This is a Value of type String, '%s'", inp)
	case bool:
		result = fmt.Sprintf("This is a Value of type Boolean, '%t'", inp)
	case Hello:
		result = fmt.Sprintf("This is a Value of type Hello, '%s'", inp)
	default:
		result = "Enter Valid Choice !!"
	}
	return result
}

func main() {

	var no int
	fmt.Println("Enter Your Input Between 1 to 4 : ")

	_, err := fmt.Scan(&no)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if no < 1 || no > 4 {
		fmt.Println("*** Plz, Enter Valid Choice !!")
		return
	}

	var val AnyType
	switch no {
	case 1:
		val = 5
		fmt.Println(AcceptAnything(val))
	case 2:
		val = "Welcome"
		fmt.Println(AcceptAnything(val))
	case 3:
		val = true
		fmt.Println(AcceptAnything(val))
	case 4:
		var myVar Hello
		myVar = "My Data"
		fmt.Println(AcceptAnything(myVar))
	default:
		fmt.Println("Enter Valid Choice !!")
	}
}
