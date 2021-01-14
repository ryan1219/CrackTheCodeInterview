package binarytree

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const MARKER = "#"

type Codec struct {
	Val string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.Val = ""
	this.preorder(root)
	return this.Val
}

func (this *Codec) preorder(root *TreeNode) {
	if root == nil {
		this.Val = this.Val + MARKER + " "
		return
	}

	this.Val = this.Val + strconv.Itoa(root.Val) + " "
	this.preorder(root.Left)
	this.preorder(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	pre := strings.Split(strings.Trim(data, " "), " ")
	if len(pre) == 0 {
		return nil
	}

	node := &TreeNode{}
	root := &node
	this.preorderTraversal(root, &pre)
	return *root
}

func (this *Codec) preorderTraversal(root **TreeNode, pre *[]string) {
	if len(*pre) == 0 || (*pre)[0] == MARKER {
		*pre = (*pre)[1:]
		return
	}
	num, _ := strconv.Atoi((*pre)[0])
	*root = &TreeNode{Val: num, Left: nil, Right: nil}
	*pre = (*pre)[1:]
	this.preorderTraversal(&(*root).Left, pre)
	this.preorderTraversal(&(*root).Right, pre)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
