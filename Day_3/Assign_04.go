package main

import "fmt"

func main() {

	arr := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	var ind1, ind2 int
	fmt.Print("Enter index1 : ")
	fmt.Scanln(&ind1)

	fmt.Print("Enter index2 : ")
	fmt.Scanln(&ind2)

	if ind1 < 0 || ind2 < 0 || ind2 >= len(arr) || ind1 >= len(arr) {
		fmt.Println("Incorrect Indexes")
	} else {

		var slice1, slice2, slice3 []string
		slice1 = arr[:ind1+1]
		slice2 = arr[ind1 : ind2+1]
		slice3 = arr[ind2:]

		fmt.Println(slice1)
		fmt.Println(slice2)
		fmt.Println(slice3)
	}
}
