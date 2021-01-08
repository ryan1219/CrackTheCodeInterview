package main

// question: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]TreeNode, 0)
	queue = append(queue, *root)
	right := true

	for len(queue) > 0 {
		currentNodes := len(queue)
		nextLevel := make([]TreeNode, 0)
		currentLevel := make([]int, 0)
		for currentNodes > 0 {
			node := queue[currentNodes-1]
			currentLevel = append(currentLevel, node.Val)
			if right {
				if node.Left != nil {
					nextLevel = append(nextLevel, *node.Left)
				}
				if node.Right != nil {
					nextLevel = append(nextLevel, *node.Right)
				}
			} else {
				if node.Right != nil {
					nextLevel = append(nextLevel, *node.Right)
				}
				if node.Left != nil {
					nextLevel = append(nextLevel, *node.Left)
				}
			}
			currentNodes--
		}
		//switch direction, next level
		result = append(result, currentLevel)
		right = !right
		queue = nextLevel
	}

	return result
}
