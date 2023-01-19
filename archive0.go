package main

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

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
