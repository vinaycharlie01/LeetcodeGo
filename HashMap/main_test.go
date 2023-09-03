package main

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	type args struct {
		arr []int
	}
	a := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testcase1",
			args: args{
				arr: []int{1, 1},
			},
			want: []int{2},
		},
	}
	for _, v := range a {
		t.Run(v.name, func(t *testing.T) {
			if got := FindDisappearedNumbers(v.args.arr); !reflect.DeepEqual(got, v.want) {
				t.Errorf(" findDisappearedNumbers(), %v  = %v, want %v", v.name, got, v.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThirdMax(tt.args.nums); got != tt.want {
				t.Errorf("ThirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
