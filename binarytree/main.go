package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	InOrder(root.Right)
}

func insert(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return &TreeNode{v, nil, nil}
	}
	if v < root.Val {
		root.Left = insert(root.Left, v)
		return root
	}
	root.Right = insert(root.Right, v)
	return root
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func heightofTreee(root *TreeNode) int {
	if root == nil {
		return -1
	} else {
		b := max(heightofTreee(root.Left)+1, heightofTreee(root.Right)+1)
		return int(b)
	}
}

func Not(root *TreeNode) bool {
	if root != nil {
		return true
	}
	return false
}

func CountTreeNode(root *TreeNode) int {
	if Not(root) {
		return 1 + CountTreeNode(root.Left) + CountTreeNode(root.Right)
	} else {
		return 0
	}
}

func CountTreeNodeLeft(root *TreeNode) int {
	if Not(root) {
		return 1 + CountTreeNodeLeft(root.Left)
	}
	return 0
}
func CountTreeNodeRight(root *TreeNode) int {
	if Not(root) {
		return 1 + CountTreeNodeRight(root.Left)
	}
	return 0
}

func NewTreeNode() *TreeNode {
	var root *TreeNode
	val := []int{10, 1, 15, 3, 7, 12, 18}

	/*              10
					/ \
	               1   15

				3   7   12 18


	*/
	for _, v := range val {
		root = insert(root, v)
	}
	return root
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	// fmt.Println(nums1[:m])
	fmt.Println(nums1)

}

func removeElement(nums []int, val int) int {
	b := slices.ContainsFunc(nums, func(j int) bool {
		return nums[j] == val
	})

	// nums =slices.DeleteFunc(nums, func(j int) bool {
	// 	return nums[j] == val
	// })
	// slices.DeleteFunc()
}

func maxPoints(points [][]int) int {
	var a = math.MaxInt
	for _, v := range points {
		if v[1] > a {
			a = v[1]
		}
	}
	return a
}

func main() {
	merge([]int{1, 2, 3, 4, 0, 0, 0}, 4, []int{1, 2, 3}, 3)

}
