package main

import (
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func main() {
	root := &TreeNode{HasToy: true,
		Left: &TreeNode{HasToy: true,
			Left:  &TreeNode{HasToy: true},
			Right: &TreeNode{HasToy: false}},
		Right: &TreeNode{HasToy: false,
			Left:  &TreeNode{HasToy: true},
			Right: &TreeNode{HasToy: true}},
	}
	fmt.Println(unrollGarland(root))
}

func unrollGarland(root *TreeNode) []bool {
	queue := []*TreeNode{root}
	var res []bool
	var zig bool
	for len(queue) > 0 {
		qlen := len(queue)
		var level []bool
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.HasToy)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if zig == false {
			level = Reverse(level)
			zig = true
		} else {
			zig = false
		}
		res = append(res, level...)
	}
	return res
}

func Reverse(s []bool) []bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
