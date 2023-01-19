package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func permString(str []rune, left int, right int) {
	if left == right {
		fmt.Println(string(str)) // outputs permutation
	} else {
		for i := left; i <= right; i++ {
			str[left], str[i] = str[i], str[left]
			permString(str, left+1, right)
			str[left], str[i] = str[i], str[left]
		}
	}
}

func stringPermutations() {
	str := "ghkw"
	permString([]rune(str), 0, len(str)-1)
}
