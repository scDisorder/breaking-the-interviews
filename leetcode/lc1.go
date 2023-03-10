package leetcode

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]

	for i := range strs {
		for strings.Index(strs[i], res) != 0 {
			res = res[:len(res)-1]
			if res == "" {
				return ""
			}
		}
	}

	return res
}

// Returns true if the intervals a and b have a common element.
func doesIntervalsOverlap(a, b []int) bool {
	return minInt(a[1], b[1])-maxInt(a[0], b[0]) >= 0
}

// Return the interval having all the elements of intervals a and b.
func mergeIntervals(a, b []int) []int {
	return []int{minInt(a[0], b[0]), maxInt(a[1], b[1])}
}

// https://leetcode.com/problems/insert-interval/
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i := 0
	n := len(intervals)
	start, end := newInterval[0], newInterval[1]
	for i < n && intervals[i][1] < start {
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= end {
		start = minInt(start, intervals[i][0])
		end = maxInt(end, intervals[i][1])
		i++
	}
	res = append(res, []int{start, end})
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	dfsInts(nums, 0, []int{}, &res)
	return res
}

func dfsInts(nums []int, start int, cur []int, res *[][]int) {
	if len(cur) >= 2 {
		*res = append(*res, append([]int{}, cur...))
	}

	m := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}

		if len(cur) == 0 || nums[i] >= cur[len(cur)-1] {
			m[nums[i]] = true
			dfsInts(nums, i+1, append(cur, nums[i]), res)
		}
	}
}
