package main

import (
	"fmt"
)

func message(alice chan string, bob chan string, flag chan bool) {

	for {
		select {
		case val := <-alice:
			fmt.Printf("alice: %s,", val)

		case val := <-bob:
			fmt.Printf("bob: %s,", val)

		case <-flag:
			return
		}
	}
}

func main() {
	s := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	alice := make(chan string)
	bob := make(chan string)
	flag := make(chan bool)

	go message(alice, bob, flag)

	var temp string
	for i := 0; i < len(s); i++ {
		ch := fmt.Sprintf("%c", s[i])

		if ch == "$" {
			alice <- temp
			temp = ""
		} else if ch == "#" {
			bob <- temp
			temp = ""
		} else if ch == "^" {
			flag <- true
			break
		} else {
			temp = temp + ch
		}
	}
}
