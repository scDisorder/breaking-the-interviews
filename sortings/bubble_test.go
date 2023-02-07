package sortings

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		src []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-01",
			args: args{src: []int{2, 9, 7, 6, 3, 1, 0, 8}},
			want: arr0Sorted,
		},
		{
			name: "test-02",
			args: args{src: []int{64, 34, 25, 12, 22, 11, 90}},
			want: arr1Sorted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.src)
			if !reflect.DeepEqual(tt.want, tt.args.src) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}
}

func TestCocktailSort(t *testing.T) {
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
			CocktailSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestCombSort(t *testing.T) {
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
			CombSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
