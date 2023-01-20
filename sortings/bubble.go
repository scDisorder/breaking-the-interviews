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
