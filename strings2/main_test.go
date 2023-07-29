package main

import (
	"reflect"
	"testing"
)

type SplitWordInput struct {
	Words     []string
	Separator byte
}
type SplitWordsBySeparatorData struct {
	Testcase string
	Input    SplitWordInput
	Result   []string
}

func TestSplitWordsBySeparator(t *testing.T) {
	run := []SplitWordsBySeparatorData{
		{
			Testcase: "testcase1",
			Input: SplitWordInput{
				Words:     []string{"one.two.three", "four.five", "six"},
				Separator: '.',
			},
			Result: []string{"one", "two", "three", "four", "five", "six"},
		},
		{
			Testcase: "testcase2",
			Input: SplitWordInput{
				Words:     []string{"$easy$", "$problem$"},
				Separator: '$',
			},
			Result: []string{"easy", "problem"},
		},
		// {
		// 	Testcase: "testcase3",
		// 	Input: SplitWordInput{
		// 		Words:     []string{"|||"},
		// 		Separator: '|',
		// 	},
		// 	Result: []string{},
		// },
	}
	for _, v := range run {
		t.Run(v.Testcase, func(t *testing.T) {
			got := SplitWordsBySeparator(v.Input.Words, v.Input.Separator)
			// fmt.Println(got, v.Result)
			if !reflect.DeepEqual(got, v.Result) {
				t.Errorf("%v , got %v, want %v", v.Testcase, got, v.Result)
			}
		})
	}
}
