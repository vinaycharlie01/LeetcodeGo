package main

import "testing"

type Testcases struct {
	cases string
	data  []int
	want  int
}

func TestThirdMax(t *testing.T) {
	chek := []Testcases{{
		cases: "test1",
		data:  []int{3, 2, 1},
		want:  1,
	}, {
		cases: "test2",
		data:  []int{1, 2},
		want:  2,
	}, {
		cases: "test3",
		data:  []int{2, 2, 3, 1},
		want:  3,
	}}
	for _, v := range chek {
		result := thirdMax(v.data)
		if v.want != result {
			t.Errorf(" %v This test case Failed expected %v, got %v", v.cases, v.want, result)
		}
	}
}
