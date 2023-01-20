package sortings

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
