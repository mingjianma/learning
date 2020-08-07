package main

import "fmt"

//给定两个二叉树，编写一个函数来检验它们是否相同。
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	b := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(a, b))

	a = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	b = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(a, b))

	a = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	b = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(a, b))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//两棵树都是nil时返回true 否则返回false
	if p == nil || q == nil {
		if q == p {
			return true
		} else {
			return false
		}
	}
	//两个节点的val不相等时返回false
	if p.Val != q.Val {
		return false
	}

	//递归检查两个节点左边的子节点，当子节点有差异，返回false
	if p.Left != nil || q.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}
	//递归检查两个节点右边边的子节点，当子节点有差异，返回false
	if p.Right != nil || q.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}
	//都通过表示两个节点一致
	return true
}
