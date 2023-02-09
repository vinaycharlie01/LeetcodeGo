package main

import (
	"reflect"
	"testing"
)

type Testcases struct {
	cases string
	user  string
	dp    map[string]int
	want  map[string]int
}

func TestRecone(t *testing.T) {
	hello := []Testcases{{
		cases: "1",
		user:  "vinay",
		dp:    map[string]int{},
		want:  map[string]int{"v": 1, "i": 1, "n": 1, "a": 1, "y": 1},
	}, {cases: "2",
		user: "Bharath",
		dp:   map[string]int{},
		want: map[string]int{"B": 1, "h": 2, "a": 2, "r": 1, "t": 1},
	},
		{cases: "3",
			user: "Charlie",
			dp:   map[string]int{},
			want: map[string]int{"C": 1, "h": 1, "a": 1, "r": 1, "i": 1, "e": 1, "l": 1},
		},
	}

	for _, v := range hello {
		result := Recone(v.user, v.dp)
		if !(reflect.DeepEqual(result, v.want)) {
			t.Errorf("Test case %v: got %v, want %v", v.cases, result, v.want)
		}
	}
}
