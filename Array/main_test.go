package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersection(t *testing.T) {
	Userinput := [][][]int{{{1, 2, 2, 1}, {2, 2}}, {{4, 9, 5}, {9, 4, 9, 8, 4}}, {{1, 2, 3, 2, 4, 5, 6, 7, 8, 9}, {1}}, {{1}, {0}}}
	Output := [][]int{{2}, {4, 9}, {1}, {}}
	var a [][]int
	for _, v := range Userinput {
		Result := intersection(v[0], v[1])
		a = append(a, Result)
	}
	// fmt.Println(a)
	var map2 = map[int]int{}
	for _, v := range a {
		for _, val := range v {
			map2[val]++
		}
	}
	var map3 = map[int]int{}
	for _, v := range Output {
		for _, val := range v {
			map3[val]++
		}
	}
	require.Equal(t, map2, map3)

}

func TestContainsDuplicate(t *testing.T) {
	chek := [][]int{{1, 2, 3, 1}, {1, 2, 3, 4}, {1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}
	result := []bool{true, false, true}
	if len(chek) < 0 {
		t.Error("Chek the your data is empty", chek)
	} else if len(result) != len(chek) {
		t.Error("not enter your output", result)
	} else {
		i := 0
		for _, v := range chek {
			result2 := containsDuplicate(v)
			require.Equal(t, result[i], result2)
			i++
		}
	}
}

func TestCal(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		resut bool
	}{{
		name:  "testcase1",
		input: []int{1, 2, 3, 1},
		resut: true,
	}, {
		name:  "testcase2",
		input: []int{1, 2, 3, 4},
		resut: true,
	}}

	for _, v := range testcases {
		result := containsDuplicate(v.input)
		if result != v.resut {

		}
	}
}
