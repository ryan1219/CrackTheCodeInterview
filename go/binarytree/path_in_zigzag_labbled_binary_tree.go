package binarytree

import (
	"math"
)

// question: https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
/*
Height of node – The height of a node is the number of edges on the longest downward path between that node and a leaf.
Depth –The depth of a node is the number of edges from the node to the tree's root node.
Depth of the root is 0
Level – The level of a node is defined by 1 + the number of connections between the node and the root.
Level = Depth + 1

label on each level is 2^(level-1) < level <= 2^level - 1
*/
func pathInZigZagTree(label int) []int {
	// find which level that label stays on
	level := 0
	for label > int(math.Pow(2, float64(level)))-1 {
		level++
	}

	res := make([]int, 0)
	for label >= 1 {
		res = append(res, label)
		label = int(math.Pow(2, float64(level))) - 1 - label + int(math.Pow(2, float64(level-1)))
		label = label / 2
		level--
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
