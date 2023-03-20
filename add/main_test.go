// package main

// import (
// 	"testing"
// )

// func TestSum(t *testing.T) {
// 	// a := []int{10, 2, 4, 5, 7}
// 	// b := []int{10, 2, 4, 10, 14}

// 	result := smallestEvenMultiple(10)
// 	if result != 10 {
// 		t.Errorf("got %q, wanted %q", result, 10)
// 	}
// 	// // a := smallestEvenMultiple(10)
// 	// for i, v := range a {
// 	// 	result := smallestEvenMultiple(v)
// 	// 	require.Equal(t, result, b[i])
// 	// }

// }

package main

import "testing"

type TwoSumIP struct {
	arr    []int
	target int
}
type twoSumTC struct {
	cases string
	input TwoSumIP
	want  []int
}

func TestTwoSum(t *testing.T) {
	a := []twoSumTC{
		{
			cases: "TS1",
			input: TwoSumIP{arr: []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			cases: "TS2",
			input: TwoSumIP{arr: []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			cases: "TS1",
			input: TwoSumIP{arr: []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}

	for _, v := range a {
		result := TwoSum(v.input.arr, v.input.target)
		for i := range result {
			if result[i] != v.want[i] {
				t.Errorf("%s Testcase not accepted", v.cases)
			}
		}

	}
}

type testcasesd struct {
	name string
	a    int
	b    int
	want int
}

func TestAdd(t *testing.T) {
	a := []testcasesd{
		{
			name: "tescase1",
			a:    10,
			b:    20,
			want: 30,
		}, {
			name: "tescase2",
			a:    1,
			b:    2,
			want: 2,
		}, {
			name: "tescase3",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, v := range a {
		res := ADD(v.a, v.b)
		if res != v.want {
			t.Errorf("%v", v.name)
		}
	}
}
