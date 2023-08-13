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

func TestFinalString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantA string
	}{
		// TODO: Add test cases.
		{
			name:  "test1",
			args:  args{s: "string"},
			wantA: "rtsng",
		},
		{
			name:  "test2",
			args:  args{s: "poiinter"},
			wantA: "ponter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := FinalString(tt.args.s); gotA != tt.wantA {
				t.Errorf("FinalString() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func Test_ReverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: struct {
				s string
				k int
			}{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "Example 2",
			args: struct {
				s string
				k int
			}{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
