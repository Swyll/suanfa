package main

import "fmt"

/**
684. 冗余连接

树可以看成是一个连通且 无环 的 无向 图。
给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，
且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。

[[1,4],[3,4],[1,3],[1,2],[4,5]]
[[1,3],[3,4],[1,5],[3,5],[2,3]]
**/
func main() {
	//fmt.Println(findRedundantConnection([][]int{[]int{1, 4}, []int{3, 4}, []int{1, 3}, []int{1, 2}, []int{4, 5}}))
	fmt.Println(findRedundantConnection([][]int{[]int{1, 3}, []int{3, 4}, []int{1, 5}, []int{3, 5}, []int{2, 3}}))
}

//并查集，每个集合有个代表节点，如果两个节点所在的集合集合代表节点相同，说明已经联通了，那么再加边就成环了
func findRedundantConnection(edges [][]int) []int {
	//初始化每个集合代表节点，初始化的时候，一个节点就是一个集合代表节点就是本身
	cache := make([]int, len(edges)+1)

	for i, _ := range cache {
		cache[i] = i
	}

	findparent := func(index int) int {
		for cache[index] != index {
			index = cache[index]
		}

		return index
	}
	for _, edge := range edges {
		//查找两节点所在集合的集合代表节点
		p1 := findparent(edge[0])
		p2 := findparent(edge[1])

		if p1 == p2 {
			return edge
		} else {
			//合并两集合,合并的时候把代表节点的代表节点修改就行，不能直接修改自身代表节点,不然代表链路会丢失
			cache[p1] = p2
		}
	}

	return []int{}
}
