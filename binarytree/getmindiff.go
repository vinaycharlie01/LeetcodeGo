package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		InOrdergetmin(root, ch)
	}()
	mini := 1<<63 - 1
	prev := -1
	flag := false
	for v := range ch {
		if !flag {
			mini = min(mini, int(math.Abs(float64(v-prev))))
		}
		prev = v
		flag = true
	}
	return mini
}

func InOrdergetmin(root *TreeNode, ch chan int) {
	if root == nil {
		return
	}
	InOrdergetmin(root.Left, ch)
	ch <- root.Val
	InOrdergetmin(root.Right, ch)

}
