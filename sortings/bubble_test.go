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
			want: []int{0, 1, 2, 3, 6, 7, 8, 9},
		},
		{
			name: "test-02",
			args: args{src: []int{64, 34, 25, 12, 22, 11, 90}},
			want: []int{11, 12, 22, 25, 34, 64, 90},
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
