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
