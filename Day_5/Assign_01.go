package main

import "fmt"

func accessSlice(slice []int, index int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Internal Error: %v\n", r)
		}
	}()

	if index >= len(slice) {
		panic("Index out of range")
	}

	if index < 0 {
		panic("Index cannot be negative")
	}

	fmt.Printf("item %d, value %d", index, slice[index])

}

func main() {

	var no, value int

	fmt.Println("Enter the size of slice")
	fmt.Scan(&no)
	slice := make([]int, no)

	fmt.Println("Enter Slice Values")
	for i := 0; i < no; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println("Enter Value")
	fmt.Scan(&value)
	accessSlice(slice, value)
}
