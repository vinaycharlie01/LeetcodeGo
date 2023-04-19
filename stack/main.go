package main

import "strconv"

type Stackfun interface {
	Append(a int)
	Pop()
	AppendPlus()
	Top() int
	Sum() int
}

type Stack struct {
	data []int
}

func (s *Stack) Append(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}
func (s *Stack) AppendPlus() {
	s.data = append(s.data, s.data[len(s.data)-1]+s.data[len(s.data)-2])
}
func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Sum() int {
	var count int
	for _, v := range s.data {
		count += v
	}
	return count
}

func CalPoints2(nums []string) int {
	var res2 Stackfun
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] == "D":
			res2.Top()
			res2.Append(2 * res2.Top())
		case nums[i] == "C" && i > 1:
			res2.Pop()
		case nums[i] == "+" && i > 1:
			res2.AppendPlus()
		default:
			res, _ := strconv.Atoi(nums[i])
			res2.Append(res)
		}

	}
	return res2.Sum()
}

func main() {

}
