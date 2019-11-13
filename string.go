package main

import "fmt"

func main()  {
	var s = "hello"
	for _, char := range s  {
		fmt.Println(char)
		fmt.Printf("ascii: %c  %d\n", char, char)
		fmt.Printf("Unicode: %c  %d\n", char, char)
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count = make(map[int32]int)
	for _, char := range s  {
		count[char]++
	}
	for _, char := range t {
		count[char]--
	}
	for _, num := range count{
		if num != 0 {
			return false
		}
	}
	return true
}