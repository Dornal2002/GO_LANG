package main

import "fmt"

func accessSlice(slice []int, index int) (int, error) {

	if index >= len(slice) || index < 0 {
		return 0, fmt.Errorf("length of the slice should be more than index")
	}

	return slice[index], nil

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
	res1, err := accessSlice(slice, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("item %d, value %d", res1, slice[res1])

}
