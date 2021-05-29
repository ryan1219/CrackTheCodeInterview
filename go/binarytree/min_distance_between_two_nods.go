package binarytree

/*
question:
https://www.techiedelight.com/distance-between-given-pairs-of-nodes-binary-tree/
https://leetcode.com/discuss/interview-question/algorithms/125084/given-a-binary-search-tree-find-the-distance-between-2-nodes

Write a function that given a BST, it will return the distance (number of edges) between 2 nodes.

For example, given this tree

         5
        / \
       3   6
      / \   \
     2   4   7
    /         \
   1           8
The distance between 1 and 4 is 3: [1 -> 2 -> 3 -> 4]

The distance between 1 and 8 is 6: [1 -> 2 -> 3 -> 5 -> 6 -> 7 -> 8]
*/
type Node struct {
	data        int
	left, right *Node
}

func findNodesDistance(root *Node, a, b int) int {
	commonAncestor := findCommonAncestor(root, a, b)
	if commonAncestor == nil {
		return -1
	}

	distanceA := findPath(commonAncestor, a, 0)
	distanceB := findPath(commonAncestor, b, 0)

	return distanceA + distanceB
}

func findCommonAncestor(root *Node, a, b int) *Node {
	if root == nil || root.data == a || root.data == b {
		return root
	}

	left := findCommonAncestor(root.left, a, b)
	right := findCommonAncestor(root.right, a, b)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}

	return nil
}

func findPath(root *Node, target, path int) int {
	if root == nil {
		return -1
	}

	if root.data == target {
		return path
	}

	left := findPath(root.left, target, path+1)
	right := findPath(root.right, target, path+1)
	if left != -1 {
		return left
	} else if right != -1 {
		return right
	} else {
		return -1
	}
}
