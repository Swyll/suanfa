package main

import "fmt"

/**
给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。
子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。
如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。
**/
func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 0}, []int{1, 1, 1}}, 0))
}

//按前面数组的思路来讲矩阵抽象为数组，遍历下边界，右边界，上边界，左边界由hash存储
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	count := 0

	for i, v := range matrix {
		//固定了下边，遍历上边和右边
		for l := 0; l <= i; l++ {
			preCount := 0
			cacheM := map[int]int{0: 1}
			//固定了下边和上边，遍历右边
			for j, v2 := range v {
				//替换二维数组为前缀和只需要替换一次
				if l == 0 {
					preCount += v2
					if i == 0 {
						matrix[i][j] = preCount
					} else {
						matrix[i][j] = preCount + matrix[i-1][j]
					}
				}

				var pre int
				if l == 0 {
					pre = matrix[i][j]
				} else {
					pre = matrix[i][j] - matrix[l-1][j]
				}
				count += cacheM[pre-target]
				cacheM[pre] += 1
			}
		}
	}

	return count
}
