package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		row := make([]int, 0)
		temp := make([]*TreeNode, 0)
		for _, v := range queue {
			row = append(row, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		res = append(res, row)
		queue = temp
	}

	return res
}
