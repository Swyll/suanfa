package main

/**
210. 课程表 II

现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，
其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
**/

func main() {
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	//存储每个节点的下一节点集合
	cs := make([][]int, numCourses)
	//存储每个节点的入度
	cache := make([]int, numCourses)
	result := make([]int, 0, numCourses)

	//统计每个节点的下一节点
	for _, v := range prerequisites {
		cs[v[1]] = append(cs[v[1]], v[0])
		cache[v[0]]++
	}

	//存储所有入度为0的节点也就是可以作为起始点的节点
	heads := make([]int, 0, 4)
	//找出每个入度为0的节点
	for index, v := range cache {
		if v == 0 {
			heads = append(heads, index)
		}
	}

	for len(heads) != 0 {
		u := heads[0]
		//出队列
		heads = heads[1:]
		result = append(result, u)

		for _, v := range cs[u] {
			cache[v]--
			//如果入度为0说明所有依赖的节点都走完了，那么入队列
			if cache[v] == 0 {
				heads = append(heads, v)
			}
		}
	}

	//如果结果不为课程总数，说明还有节点入度没清零，那么就存在成环点
	if len(result) != numCourses {
		return []int{}
	}

	return result
}
