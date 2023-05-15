package main

import (
	"fmt"
	"testing"
)

type testParallelSaver1 struct {
	errCode int
}

func (s *testParallelSaver1) Save(v any) error {
	if intVal, ok := v.(int); ok {
		if intVal == s.errCode {
			return fmt.Errorf("int val error: %d", intVal)
		}
	}

	fmt.Println("Saved successfully", v)
	return nil
}

func Test_handleParallelTasks(t *testing.T) {
	type args struct {
		s Saver
		v []any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test-00",
			args: args{
				s: &testParallelSaver1{},
				v: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			},
			wantErr: false,
		},
		{
			name: "test-01",
			args: args{
				s: &testParallelSaver1{errCode: 2},
				v: []any{1, 3, 4, 5, 2},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleParallelTasks(tt.args.s, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("handleParallelTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
