package main

import "testing"

func TestCountSymmetricIntegers2(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				low:  1,
				high: 100,
			},
			want: 9,
		},
		{
			name: "test2",
			args: args{
				low:  1200,
				high: 1230,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSymmetricIntegers2(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("CountSymmetricIntegers2() = %v, want %v", got, tt.want)
			}
		})
	}
}
