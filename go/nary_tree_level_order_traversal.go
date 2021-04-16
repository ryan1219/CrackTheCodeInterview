package main

import "container/list"

// question: https://leetcode.com/problems/n-ary-tree-level-order-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := list.New()
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue.PushFront(root)

	for queue.Len() > 0 {
		size := queue.Len()
		level := make([]int, 0)
		for size > 0 {
			cur := queue.Back()
			queue.Remove(cur)
			level = append(level, cur.Value.(*Node).Val)
			for _, child := range cur.Value.(*Node).Children {
				queue.PushFront(child)
			}
			size--
		}
		res = append(res, level)
	}

	return res
}
