package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func main() {
	tree := &TreeNode{HasToy: false,
		Right: &TreeNode{HasToy: false,
			Right: &TreeNode{HasToy: true}},
		Left: &TreeNode{HasToy: true,
			Right: &TreeNode{HasToy: true}}}

	fmt.Println(areToysBalanced(tree))
}

func areToysBalanced(root *TreeNode) bool {
	left := counter(root.Left)
	right := counter(root.Right)
	if left+num(root.HasToy) != right+num(root.HasToy) {
		return false
	}
	return true
}

func counter(root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return num(root.HasToy)
	} else if root.Right != nil && root.Left == nil {
		return counter(root.Right) + num(root.HasToy)
	} else if root.Right == nil && root.Left != nil {
		return counter(root.Left) + num(root.HasToy)
	}
	return counter(root.Left) + counter(root.Right) + num(root.HasToy)
}

func num(hasToy bool) int {
	if hasToy == false {
		return 0
	}
	return 1
}
