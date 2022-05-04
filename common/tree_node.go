package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toString(node *TreeNode) string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("(%d, %s, %s)", node.Val, toString(node.Left), toString(node.Right))
}

func (node *TreeNode) String() string {
	return toString(node)
}

func BuildTree(nodes []any) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	root := TreeNode{nodes[0].(int), nil, nil}
	queue := []*TreeNode{&root}
	index := 1

	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}

		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]
			if index >= len(nodes) {
				cur.Left = nil
			} else if nodes[index] == nil {
				cur.Left = nil
			} else {
				cur.Left = &TreeNode{nodes[index].(int), nil, nil}
				queue = append(queue, cur.Left)
			}

			index++

			if index >= len(nodes) {
				cur.Right = nil
			} else if nodes[index] == nil {
				cur.Right = nil
			} else {
				cur.Right = &TreeNode{nodes[index].(int), nil, nil}
				queue = append(queue, cur.Right)
			}

			index++
		}
	}

	return &root
}
