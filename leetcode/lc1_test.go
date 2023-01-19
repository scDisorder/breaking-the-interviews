package leetcode

import "testing"

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
