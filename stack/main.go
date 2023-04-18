package main

import "strconv"

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
	var res2 Stack
	for i := 0; i < len(nums); i++ {
		res, err := strconv.Atoi(nums[i])
		if err != nil {
			if nums[i] == "D" {
				res2.Top()
				res2.Append(2 * res2.Top())
			}
			if nums[i] == "C" && i > 1 {
				res2.Pop()
			}
			if nums[i] == "+" && i > 1 {
				res2.AppendPlus()
			}
		} else {
			res2.Append(res)
		}
	}
	return res2.Sum()
}

func main() {

}
