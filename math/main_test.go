package main

import "testing"

type Inputtest struct {
	Case string
	A, B int
	Want int
}

func TestFindDelayedArrivalTime(t *testing.T) {
	b := []Inputtest{
		{Case: "test1", A: 15, B: 5, Want: 20},
		{Case: "test2", A: 13, B: 11, Want: 0},
	}
	for _, v := range b {
		t.Run(v.Case, func(t *testing.T) {
			res := FindDelayedArrivalTime(v.A, v.B)
			if res != v.Want {
				t.Errorf("  got %v, want %v", res, v.Want)
			}
		})
	}
}
