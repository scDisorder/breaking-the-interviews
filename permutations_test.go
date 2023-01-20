package main

import (
	"reflect"
	"testing"
)

func Test_permString(t *testing.T) {
	type args struct {
		src  string
		left int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-01",
			args: args{
				src:  "abc",
				left: 0,
			},
			want: []string{"abc", "acb", "bac", "bca", "cba", "cab"},
		},
		{
			name: "test-02",
			args: args{
				src:  "qwe",
				left: 0,
			},
			want: []string{"qwe", "qew", "wqe", "weq", "ewq", "eqw"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permString(tt.args.src, tt.args.left); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permInt(t *testing.T) {
	type args struct {
		v []int
		i int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-03",
			args: args{
				v: []int{1, 2, 3},
				i: 0,
			},
			want: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permInt(tt.args.v, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
