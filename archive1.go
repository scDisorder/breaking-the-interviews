package main

import "fmt"

func double(arr []*int) {
	for i, _ := range arr {
		val := arr[i]
		x := *val * 2
		arr[i] = &x
	}
}

func reverse(v []int) []int {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	return v
}

func sortings() {
	var zeroSlice []int
	var zeroArr [8]int

	fmt.Printf("slice: %#v\n ", zeroSlice)
	fmt.Printf("array: %#v\n", zeroArr)
	fmt.Println([]string(nil) == nil)
	fmt.Printf("%#v\n", append([]string(nil), "x"))
	fmt.Printf("%#v\n", append([]string(nil), []string(nil)...))

	for i := range zeroSlice {
		fmt.Printf("zero slice: %d\n", i)
	}

	sortFoosByBars()

}

func fillString(i int) string {
	var res string
	var printed bool
	if i%3 == 0 {
		res += "fizz"
		printed = true
	}

	if i%5 == 0 {
		if len(res) > 0 {
			res += " "
		}
		res += "buzz"
		printed = true
	}

	if !printed {
		res = fmt.Sprintf("%d", i)
	}

	return res
}
