package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	node := &Node{
		Val: 1,
		Neighbors: []*Node{
			{
				Val: 2,
				Neighbors: []*Node{
					{
						Val:       1,
						Neighbors: nil,
					},
					{
						Val: 3,
						Neighbors: []*Node{
							{
								Val:       2,
								Neighbors: nil,
							},
							{
								Val:       4,
								Neighbors: nil,
							},
						},
					},
				},
			},
			{
				Val: 4,
				Neighbors: []*Node{
					{
						Val:       1,
						Neighbors: nil,
					},
					{
						Val: 3,
						Neighbors: []*Node{
							{
								Val:       2,
								Neighbors: nil,
							},
							{
								Val:       4,
								Neighbors: nil,
							},
						},
					},
				},
			},
		},
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}

		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &Node{node.Val, []*Node{}}
		// 哈希表存储
		visited[node] = cloneNode

		// 遍历该节点的邻居并更新克隆节点的邻居列表
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(node)
}
