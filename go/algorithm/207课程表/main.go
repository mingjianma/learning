package main

import "container/list"

//你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
//在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

func main() {

}

//标准拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 判断有向图是否是有向无环图DAG
	// 计算课程的入度
	inDegree := make([]int, numCourses)

	// 构建邻接表
	adj := make(map[int]*list.List)
	length := len(prerequisites)
	for i := 0; i < length; i++ {
		inDegree[prerequisites[i][0]]++
		if adj[prerequisites[i][1]] == nil {
			l := list.New()
			l.PushBack(prerequisites[i][0])
			adj[prerequisites[i][1]] = l
		} else {
			adj[prerequisites[i][1]].PushBack(prerequisites[i][0])
		}
	}

	// 将入度为0的节点放入队列
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() != 0 {
		e := queue.Front()
		node := e.Value.(int)
		queue.Remove(e)

		// bfs
		if adj[node] != nil {
			for n := adj[node].Front(); n != nil; n = n.Next() {
				nextNode := n.Value.(int)
				inDegree[nextNode]--
				if inDegree[nextNode] == 0 {
					queue.PushBack(nextNode)
				}
			}
		}
	}

	// bfs过后，若存在节点的入度不为零，则说明存在环
	for i := 0; i < numCourses; i++ {
		if inDegree[i] != 0 {
			return false
		}
	}
	return true
}

//拓扑排序加Dfs
func canFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
