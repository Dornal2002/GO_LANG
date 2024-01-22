package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter string: ")

	input.Scan()
	var s string = input.Text()

	temp := ""

	freq := make(map[string]int, 0)

	var words []string

	for i := 0; i < len(s); i++ {
		ch := fmt.Sprintf("%c", s[i])

		if ch == " " && temp != "" {
			freq[temp]++
			words = append(words, temp)
			temp = ""
		} else {
			temp = temp + ch
		}
	}

	if temp != " " {
		freq[temp]++
		words = append(words, temp)
	}

	var maxFreq int = 0
	for _, cnt := range freq {
		if cnt > maxFreq {
			maxFreq = cnt
		}
	}

	var max []string

	for i := 0; i < len(words); i++ {
		if freq[words[i]] == maxFreq {
			max = append(max, words[i])
			freq[words[i]] = -1
		}
	}

	fmt.Println(max)

}
