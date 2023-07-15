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
