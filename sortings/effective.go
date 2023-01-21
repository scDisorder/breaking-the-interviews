package sortings

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSort is a divide-and-conquer algorithm that recursively partitions an array into
// two smaller arrays until each partition consists of only one element
// complexity from O(n*log(n)) to O(n^2)
func QuickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func merge(arr []int, left, middle, right int) {
	n1 := middle - left + 1
	n2 := right - middle

	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[middle+1+j]
	}

	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// MergeSort is a divide-and-conquer algorithm that recursively divides an array
// into two smaller arrays until each array consists of only one element.
// complexity O(n*log(n))
func MergeSort(arr []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		MergeSort(arr, left, middle)
		MergeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		heapify(arr, n, largest)
	}
}

func HeapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]

		heapify(arr, i, 0)
	}
}
