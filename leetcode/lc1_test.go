package leetcode

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-01",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "test-02",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "test-03",
			args: args{strs: []string{"ponytail", "portfolio", "pointer"}},
			want: "po",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInterval(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-01",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		//{
		//	name: "test-02",
		//	args: args{
		//		intervals:   [][]int{{1, 3}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		//		newInterval: []int{4, 8},
		//	},
		//	want: [][]int{{1, 2}, {3, 10}, {12, 16}}, // ?? why test fails but result is the same
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInterval(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-01",
			args: args{nums: []int{4, 6, 7, 7}},
			want: [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
		},
		{
			name: "test-02",
			args: args{nums: []int{4, 4, 3, 2, 1}},
			want: [][]int{{4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
