package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
	}
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	// return reflect.DeepEqual(p, q)
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}

func countNodes(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
	return 0
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsSameTree(root.Left, root.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	var root *TreeNode
	for v := range inorder {
		root = insert(root, v)
	}
	return root
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
