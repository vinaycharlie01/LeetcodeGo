package main

import "testing"

type testcases struct {
	cases string
	input []int
	want  int64
}

func TestString(t *testing.T) {
	a := []testcases{
		{
			cases: "testcase1",
			input: []int{7, 52, 2, 4},
			want:  596,
		},
		{

			cases: "testcase2",
			input: []int{5, 14, 13, 8, 12},
			want:  192,
		},
	}

	for _, v := range a {
		result := findTheArrayConcVal(v.input)
		if result != v.want {
			t.Errorf("%v this case will be error", v.cases)
		}

	}

}
