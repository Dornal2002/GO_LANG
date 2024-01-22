package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter the number for an Equivalent Day:")
	fmt.Scan(&no)
	day := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	res, ok := day[no]

	if !ok {
		fmt.Println("Not a day")
	} else {
		fmt.Println("Equivalent day:", res)
	}

}
