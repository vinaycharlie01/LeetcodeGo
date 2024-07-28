package main

/*
func kthSmallest(root *TreeNode, k int) int {
   ch:=make(chan int)
   b:=0
   var a *int =&b
   go PreOrder(root, k, ch, a)
   return <-ch
}

func PreOrder(root *TreeNode, k int, ch chan int, a *int) {
	if root == nil {
		return
	}
	PreOrder(root.Left, k, ch, a)
	*a++
	if *a == k {
		ch <- root.Val
		defer close(ch)
		return
	}
	PreOrder(root.Right, k, ch, a)
}

*/

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	PreOrderKth(root, &k, &res)
	return res
}

func PreOrderKth(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}
	PreOrderKth(root.Left, k, res)
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}
	PreOrderKth(root.Right, k, res)
}
