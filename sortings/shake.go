package sortings

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
