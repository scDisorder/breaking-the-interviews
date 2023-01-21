package sortings

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr1 := []int{2, 9, 7, 6, 3, 1, 0, 8}
	arr2 := []int{64, 34, 25, 12, 22, 11, 90}

	type args struct {
		arr  []int
		low  int
		high int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-01",
			args: args{
				arr:  arr1,
				low:  0,
				high: len(arr1) - 1,
			},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{
				arr:  arr2,
				low:  0,
				high: len(arr2) - 1,
			},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	arr1 := []int{2, 9, 7, 6, 3, 1, 0, 8}
	arr2 := []int{64, 34, 25, 12, 22, 11, 90}

	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-01",
			args: args{
				arr:   arr1,
				left:  0,
				right: len(arr1) - 1,
			},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{
				arr:   arr2,
				left:  0,
				right: len(arr2) - 1,
			},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr, tt.args.left, tt.args.right)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	arr1 := []int{2, 9, 7, 6, 3, 1, 0, 8}
	arr2 := []int{64, 34, 25, 12, 22, 11, 90}

	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-01",
			args: args{
				arr: arr1,
			},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{
				arr: arr2,
			},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
