package sortings

// from O(n) to O(n^2)
func BubbleSort(src []int) {
	n := len(src)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
}

// O(n^2)
func CocktailSort(arr []int) {
	n := len(arr)
	left := 0
	right := n - 1
	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left++
	}
}

func CombSort(arr []int) {
	n := len(arr)
	gap := n
	shrink := 1.3
	sorted := false
	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap > 1 {
			sorted = false
		} else {
			gap = 1
			sorted = true
		}
		i := 0
		for i+gap < n {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				sorted = false
			}
			i++
		}
	}
}
