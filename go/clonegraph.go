package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// question: https://leetcode.com/problems/clone-graph/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	seen := make(map[*Node]*Node)
	next := make([]*Node, 0)
	copy := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	seen[node] = copy
	next = append(next, node)

	for len(next) != 0 {
		tmp := next[0]
		if len(seen[tmp].Neighbors) == 0 {
			for _, i := range tmp.Neighbors {
				if _, found := seen[i]; found {
					seen[tmp].Neighbors = append(seen[tmp].Neighbors, seen[i])
				} else {
					iCopy := &Node{Val: i.Val, Neighbors: make([]*Node, 0)}
					seen[tmp].Neighbors = append(seen[tmp].Neighbors, iCopy)
					seen[i] = iCopy
					next = append(next, i)
				}
			}
		}

		next = next[1:]
	}

	return copy
}
