package main

import "sort"

// question: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	m := make(map[int][]nodePosition)
	res := make([][]int, 0)
	stack := make([]nodePosition, 0)
	stack = append(stack, nodePosition{root, 0, 0})

	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if _, contained := m[cur.col]; !contained {
			m[cur.col] = make([]nodePosition, 0)
		}
		m[cur.col] = append(m[cur.col], cur)
		if cur.node.Left != nil {
			stack = append(stack, nodePosition{cur.node.Left, cur.row + 1, cur.col - 1})
		}
		if cur.node.Right != nil {
			stack = append(stack, nodePosition{cur.node.Right, cur.row + 1, cur.col + 1})
		}
	}

	// sort the map by keys
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		list := m[k]
		sort.Slice(list, func(i, j int) bool {
			if list[i].row == list[j].row {
				return list[i].node.Val < list[j].node.Val
			}
			return list[i].row < list[j].row
		})
		values := make([]int, 0)
		for _, v := range list {
			values = append(values, v.node.Val)
		}
		res = append(res, values)
	}

	return res
}

type nodePosition struct {
	node *TreeNode
	row  int
	col  int
}
