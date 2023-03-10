package sortings

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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
			args: args{arr: []int{2, 9, 7, 6, 3, 1, 0, 8}},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{arr: []int{64, 34, 25, 12, 22, 11, 90}},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
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
			args: args{arr: []int{2, 9, 7, 6, 3, 1, 0, 8}},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{arr: []int{64, 34, 25, 12, 22, 11, 90}},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
