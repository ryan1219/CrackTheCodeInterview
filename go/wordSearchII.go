package main

type trieNode struct {
	Next []*trieNode
	Word string
}

func newTrieNode() *trieNode {
	return &trieNode{Next: make([]*trieNode, 26), Word: ""}
}

func buildTrie(words []string) *trieNode {
	root := newTrieNode()
	for _, w := range words {
		p := root
		for _, c := range w {
			i := c - 'a'
			if p.Next[i] == nil {
				p.Next[i] = newTrieNode()
			}
			p = p.Next[i]
		}
		p.Word = w
	}

	return root
}

func dfs(board [][]byte, i, j int, root *trieNode, res *[]string) {
	c := board[i][j]

	// '#' is previous traversed
	if c == '#' || root.Next[c-'a'] == nil {
		return
	}
	root = root.Next[c-'a']
	if root.Word != "" {
		(*res) = append((*res), root.Word)
		root.Word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, root, res)
	}
	if j > 0 {
		dfs(board, i, j-1, root, res)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, root, res)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, root, res)
	}
	board[i][j] = c
}

func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0)
	root := buildTrie(words)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, i, j, root, &result)
		}
	}

	return result
}
