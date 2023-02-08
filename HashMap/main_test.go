package main

import "testing"

func TestHash(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				[]int{1, 2, 3, 4, 4, 5, 6},
				3,
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hash(tt.args.arr, tt.args.target)
		})
	}
}
