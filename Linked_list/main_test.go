package main

import (
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
