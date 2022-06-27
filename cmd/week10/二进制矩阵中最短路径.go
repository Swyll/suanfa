package main

import (
	"math"
)

/**
1091. 二进制矩阵中的最短路径

给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。

二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：

路径途经的所有单元格都的值都是 0 。
路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
畅通路径的长度 是该路径途经的单元格总数。

输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
输出：4
**/

func main() {
}

var dx = []int{1, 1, 1, 0, 0, -1, -1, -1}
var dy = []int{1, 0, -1, 1, -1, 1, 0, -1}

type node struct {
	x, y, l int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	vis := make([][]int, len(grid))
	for i := range vis {
		vis[i] = make([]int, len(grid[i]))
		for j := range vis[i] {
			vis[i][j] = math.MaxInt64
		}
	}
	vis[0][0] = 1
	queue := []node{{0, 0, 1}}
	for len(queue) > 0 {
		l := len(queue)
		for _, q := range queue {
			for i := range dx {
				x := q.x + dx[i]
				y := q.y + dy[i]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 1 {
					continue
				}
				if vis[x][y] > q.l+1 {
					vis[x][y] = q.l + 1
					queue = append(queue, node{x, y, q.l + 1})
				}
			}
		}
		queue = queue[l:]
	}
	if vis[len(grid)-1][len(grid[0])-1] == math.MaxInt64 {
		return -1
	}
	return vis[len(grid)-1][len(grid[0])-1]
}
