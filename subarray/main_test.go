package main

import (
	"reflect"
	"testing"
)

type Data struct {
	Result string
	Input  []int
	Output [][]int
}

func TestSubsets(t *testing.T) {

	Test := []Data{
		{
			Result: "test1",
			Input:  []int{1, 2, 3},
			Output: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, v := range Test {
		t.Run(v.Result, func(t *testing.T) {
			if got := Subsets(v.Input); reflect.DeepEqual(got, v.Output) {
				t.Errorf("%v , got %v, want %v", v.Result, got, v.Output)
			}
		})
		// res1 := Subsets(v.Input)
		// if !(reflect.DeepEqual(res1, v.Output)) {
		// 	t.Error(v.Result, "this case failed")
		// }

	}
}

func TestCountCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 3, 1, 2, 2}}, 4},
		{"test2", args{[]int{5, 5, 5, 5}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("CountCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
