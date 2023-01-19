package main

import "fmt"

// Perm prints out int array permutations
func Perm(v []int, f func([]int)) {
	perm(v, f, 0)
}

func perm(v []int, f func([]int), i int) {
	if i > len(v) {
		f(v)
		return
	}
	perm(v, f, i+1)
	for j := i + 1; j < len(v); j++ {
		v[i], v[j] = v[j], v[i]
		perm(v, f, i+1)
		v[i], v[j] = v[j], v[i]
	}
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

// returns all the permutations of "qwer"
func stringPermutations() {
	str := "qwer"
	permString([]rune(str), 0, len(str)-1)
}
