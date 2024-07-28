package main

func isValidBST(root *TreeNode, ch chan bool) {
	go valid(root, -1<<63, 1<<63-1, ch)
}

func valid(root *TreeNode, min int, max int, ch chan bool) {
	if root == nil {
		ch <- true
		return
	}
	if root.Val <= min || root.Val >= max {
		ch <- false
		return
	}
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go valid(root.Left, min, root.Val, ch2)
	go valid(root.Right, root.Val, max, ch3)
	ch <- <-ch2 && <-ch3
}
