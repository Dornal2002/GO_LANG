package main

import (
	"fmt"
	"sync"
)

var wt sync.WaitGroup

func reverseString(s string) {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}

	fmt.Println(res)
	wt.Done()
}

func main() {
	var s string
	fmt.Println("Enter tehe string")
	fmt.Scan((&s))
	wt.Add(1)
	go reverseString(s)
	wt.Wait()
}
