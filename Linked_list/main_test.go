package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

type userinput struct {
	data []int
	k    int
}

type Want struct {
	data []int
}

type Addtwonum struct {
	num1 []int
	num2 []int
	want []int
}

func Test_partition(t *testing.T) {

	Input := []userinput{{
		data: []int{1, 4, 3, 2, 5, 2},
		k:    3,
	}, {data: []int{2, 1},
		k: 2,
	}, {data: []int{1, 4, 3, 0, 2, 5, 2},
		k: 3,
	}, {data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		k: 4,
	},
	}
	want := [][]int{{1, 2, 2, 4, 3, 5}, {1, 2}, {1, 0, 2, 2, 4, 3, 5}, {1, 2, 3, 4, 5, 6, 7, 8, 9}}
	i := 0
	for _, s := range Input {
		Link := new(ListNode)
		temp := Link
		for _, v := range s.data {
			temp.Next = &ListNode{v, nil}
			temp = temp.Next
		}
		Result := Display(partition(Link.Next, s.k))
		require.Equal(t, Result, want[i])
		i++
	}
	// a := [][]int{{1, 4, 3, 2, 5, 2}, {2, 1}}
	// var b [][]int
	// for _, v := range a {
	// 	Link := new(ListNode)
	// 	temp := Link
	// 	for _, v := range v {
	// 		temp.Next = &ListNode{v, nil}
	// 		temp = temp.Next
	// 	}
	// 	k := 2
	// 	result := Display(partition(Link.Next, k))
	// 	b = append(b, result)
	// }
	// var c = [][]int{{1, 2, 2, 4, 3, 5}, {1, 2}}
	// i := 0
	// for _, v := range b {
	// 	require.Equal(t, c[i], v)
	// 	i++
	// }

	// type args struct {
	// 	head *ListNode
	// 	x    int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want *ListNode
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("partition() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }

}

func TestAddTwoNum(t *testing.T) {
	a := []Addtwonum{
		{
			num1: []int{2, 4, 3},
			num2: []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
	}
	for _, v := range a {
		Link := new(ListNode)
		temp := Link
		for _, v := range v.num1 {
			temp.Next = &ListNode{v, nil}
			temp = temp.Next
		}
		Link2 := new(ListNode)
		temp2 := Link2
		for _, v := range v.num2 {
			temp2.Next = &ListNode{v, nil}
			temp2 = temp2.Next
		}
		res := addTwoNumbers(Link.Next, Link2.Next)
		result := Display(res)
		fmt.Println(result)
		if !reflect.DeepEqual(result, v.want) {
			t.Error()
		}
	}
}

type list1 struct {
	name string
	a    []int
	want []int
}

func TestDelete(t *testing.T) {
	b := []list1{
		{
			name: "testcase1",
			a:    []int{1, 1, 2},
			want: []int{1, 2},
		},
	}
	for _, v := range b {
		Link := new(ListNode)
		temp := Link
		for _, v := range v.a {
			// [0:50]-[1:100]--[1:200 ]--[val:]
			newNode := &ListNode{Val: v}
			temp.Next = newNode
			temp = temp.Next
		}
		res := deleteDuplicates(Link)
		result := Display(res)
		fmt.Println(result)
		if !reflect.DeepEqual(result, v.want) {
			t.Errorf("%v", v.name)
		}
	}
}
