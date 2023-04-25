package main

import (
	"reflect"
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

type dd struct {
	data []int
	k    int
}

type Test struct {
	cases     string
	userinput dd
	want      int
}

func TestFindKthLargest(t *testing.T) {
	a := []Test{{
		cases:     "test1",
		userinput: dd{[]int{3, 2, 1, 5, 6, 4}, 2},
		want:      5,
	}, {
		cases:     "test2",
		userinput: dd{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
		want:      4,
	}, {
		cases:     "test3",
		userinput: dd{[]int{1, 2, 3, 4, 9, 7, 5, 3}, 3},
		want:      5,
	}, {
		cases:     "test4",
		userinput: dd{[]int{1, 2, 3, 4, 9, 7, 5, 3}, 6},
		want:      3,
	}, {
		cases:     "test5",
		userinput: dd{[]int{1, 2, 3, 4, 9, 7, 5, 3}, 2},
		want:      7,
	}, {
		cases:     "test6",
		userinput: dd{[]int{1}, 1},
		want:      1,
	}}
	for _, v := range a {
		result := findKthLargest(v.userinput.data, v.userinput.k)
		if result != v.want {
			t.Errorf("Error on %v: got %v, want %v", v.cases, result, v.want)
		}
	}
}

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
			want:  673,
		},
	}

	for _, v := range a {
		result := findTheArrayConcVal(v.input)
		if result != v.want {
			t.Errorf("%v this case will be error", v.cases)
		}

	}

}

// type remo struct {
// 	nums []int
// 	val  int

// }

type testcase struct {

	//testcase name
	cases string
	//input
	nums []int
	val  int

	//output
	want int
}

func TestRemoveElement(t *testing.T) {
	a := []testcase{
		{
			cases: "testcase2",
			nums:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:   2,
			want:  2,
		}, {
			cases: "testcase2",
			nums:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:   2,
			want:  2,
		},
	}
	for _, v := range a {
		result := removeElement(v.nums, v.val)
		if result != v.want {
			t.Errorf("%v", v.cases)
		}
	}
}
func BenchmarkXxx(b *testing.B) {

}

type twod struct {
	num1 [][]int
	num2 [][]int
}

type testcasesM struct {
	name  string
	input twod
	want  [][]int
}

func TestMergeArrays(t *testing.T) {
	a := []testcasesM{
		{
			name: "test1",
			input: twod{
				num1: [][]int{{1, 2}, {2, 3}, {4, 5}},
				num2: [][]int{{1, 4}, {3, 2}, {4, 1}},
			},
			want: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			name: "test2",
			input: twod{
				num1: [][]int{{2, 4}, {3, 6}, {5, 5}},
				num2: [][]int{{1, 4}, {3, 2}, {4, 1}},
			},
			want: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
	}
	for _, v := range a {
		t.Run(v.name, func(t *testing.T) {
			if got := mergeArrays(v.input.num1, v.input.num2); reflect.DeepEqual(got, v.want) {
				t.Errorf("%v , got %v, want %v", v.name, got, v.want)
			}
		})
	}
}

type TestCalPoint struct {
	Name string
	Data []string
	Want int
}

func TestCalPoints(t *testing.T) {
	a := []TestCalPoint{
		{"test1", []string{"5", "2", "C", "D", "+"}, 30},
		{"test2", []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{"test2", []string{"1", "C"}, 0},
	}
	for _, v := range a {
		t.Run(v.Name, func(t *testing.T) {
			result := CalPoints(v.Data)
			if result != v.Want {
				t.Errorf("  got %v, want %v", result, v.Want)
			}
		})
	}
}

type IsMonotonicinput struct {
	Name  string
	Input []int
	Want  bool
}

func TestIsMonotonic(t *testing.T) {
	Testcases := []IsMonotonicinput{
		{Name: "test1", Input: []int{1, 2, 3, 4}, Want: true},
		{Name: "test2", Input: []int{4, 3, 2, 1}, Want: true},
		{Name: "test3", Input: []int{1, 3, 2, 1}, Want: false},
	}
	for _, v := range Testcases {
		t.Run(v.Name, func(t *testing.T) {
			res := IsMonotonic(v.Input)
			if res != v.Want {
				t.Errorf("  got %v, want %v", res, v.Want)
			}
		})
	}
}

type Rotate2I struct {
	Name  string
	Input [][]int
	Want  [][]int
}

func TestRotate(t *testing.T) {

	a := []Rotate2I{
		{
			Name:  "test1",
			Input: [][]int{{}, {}, {}},
			Want:  [][]int{{}, {}, {}},
		},
	}
	for _, v := range a {
		t.Run(v.Name, func(t *testing.T) {
			Rotate2(v.Input)
			if reflect.DeepEqual(v.Input, v.Want) {
				t.Errorf("  got %v, want %v", v.Input, v.Want)
			}
		})
	}
}
