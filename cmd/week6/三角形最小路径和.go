package main

import (
	"fmt"
	"math"
)

/**
120.三角形最小路径和

给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。`
**/

func main() {
	//fmt.Println(minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{[]int{-11}}))
}

func minimumTotal(triangle [][]int) int {
	cache := make([][]int, 0, len(triangle))
	cache = append(cache, triangle[0])

	for i := 1; i < len(triangle); i++ {
		c := make([]int, 0, len(triangle[i]))
		c = append(c, cache[i-1][0]+triangle[i][0])

		for j := 1; j < len(triangle[i])-1; j++ {
			c = append(c, min(cache[i-1][j]+triangle[i][j], cache[i-1][j-1]+triangle[i][j]))
		}
		c = append(c, cache[i-1][len(cache)-1]+triangle[i][len(triangle[i])-1])

		cache = append(cache, c)
	}

	r := math.MaxInt
	for i := 0; i < len(cache[len(cache)-1]); i++ {
		r = min(r, cache[len(cache)-1][i])
	}

	return r
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
