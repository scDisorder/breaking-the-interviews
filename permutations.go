package main

// perInt returns int array permutations
func permInt(arr []int, left int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))
	return res
}

func permString(src string, left int) []string {
	var res []string
	str := []rune(src)

	right := len(str) - 1
	if left == right {
		res = append(res, string(str))
	} else {
		for i := left; i <= right; i++ {
			str[left], str[i] = str[i], str[left]
			res = append(res, permString(string(str), left+1)...)
			str[left], str[i] = str[i], str[left]
		}
	}

	return res
}
