package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定二叉树 [3,9,20,null,null,15,7]，
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(maxDepth(root))
	fmt.Println(maxDepthFor(root))
}

//递归，取左右节点间更大的值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := maxDepth(root.Left) + 1
	rightCount := maxDepth(root.Right) + 1
	if leftCount > rightCount {
		return leftCount
	} else {
		return rightCount
	}
}

//循环每层节点
func maxDepthFor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := []*TreeNode{root}
	maxDept := 0
	for len(nodeList) > 0 {
		sz := len(nodeList)
		for sz > 0 {
			node := nodeList[0]
			nodeList = nodeList[1:]
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			sz--
		}
		maxDept++
	}

	return maxDept
}
