package main

import (
	"breaking-the-interviews/leetcode"
	"fmt"
)

func main() {
	fmt.Println("echo fmt")

	nums := []int{1, 2, 1, 2}
	nums2 := []int{5, -3, 5}
	fmt.Println(leetcode.maxSubarraySumCircular(nums))
	fmt.Println(leetcode.maxSubarraySumCircular(nums2))
}
