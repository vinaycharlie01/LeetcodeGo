package random

import (
	"reflect"
	"testing"
)

func TestMakeRandomInts(t *testing.T) {
	type args struct {
		n int
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
			if got := MakeRandomInts(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRandomInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
