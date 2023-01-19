package leetcode

import (
	"math"
)

func TwoSum(nums []int, target int) []int {

	var currentIndex int
	var pairIndex int

outerLoop:
	for currentIndex < len(nums)-1 {
		for i, num := range nums {
			if currentIndex != i {
				if num+nums[currentIndex] == target {
					pairIndex = i
					break outerLoop
				}
			}
		}

		currentIndex++
	}

	return []int{currentIndex, pairIndex}
}

func LengthOfLongestSubstring(s string) int {
	var start, maxLength int
	checks := make(map[rune]int)

	for i, char := range s {
		if duplicateIndex, exists := checks[char]; exists && i >= start {
			maxLength = maxInt(maxLength, i-start)

			for j := start; j <= duplicateIndex; j++ {
				delete(checks, rune(s[j]))
			}
			start = duplicateIndex + 1
		}

		checks[char] = i
	}

	return maxInt(maxLength, len(s)-start)
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2
	start, end := 0, m

	for start <= end {
		i := (start + end) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			start = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			end = i - 1
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			return (maxLeft + minRight) / 2
		}
	}

	return 0
}

func RemoveDuplicates(nums []int) int {
	currentInsert := 1
	for i := range nums {
		if nums[currentInsert] != nums[i] {
			nums[currentInsert] = nums[i]
			currentInsert = i + 1
		}
	}
	return currentInsert
}

func MinFlipsMonoIncr(s string) int {
	var res int
	for _, char := range s {
		if char == '0' {
			res++
		}
	}

	ans := res
	for _, char := range s {
		if char == '0' {
			res = res - 1
			ans = minInt(ans, res)
		} else {
			res = res + 1
		}
	}

	return res
}

// uses Kadane's algorithm
// https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
func MaxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	rightMax := make([]int, n)

	for i, j := nums[n-1], n-2; j >= 0; j-- {
		i += nums[j]
		rightMax[j] = maxInt(rightMax[j+1], i)
	}

	sum1, sum2 := nums[0], nums[0]
	for i, j, k := 0, 0, 0; i < n; i++ {
		k = maxInt(k, 0) + nums[i]
		sum1 = maxInt(sum1, k)
		j += nums[i]
		if i+1 < n {
			sum2 = maxInt(sum2, j+rightMax[i+1])
		}
	}

	return maxInt(sum1, sum2)
}

func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}
