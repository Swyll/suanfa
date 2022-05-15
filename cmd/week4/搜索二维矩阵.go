package main

import (
	"fmt"
)

/**
74. 搜索二维矩阵

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
**/
func main() {
	fmt.Println(searchMatrix([][]int{[]int{1, 2, 3}, []int{5, 6, 7}, []int{11, 12, 23}}, 6))
}

//将二维数组拉平
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	right := m*n - 1
	left := 0

	//首先判断两边界
	if getValue(matrix, left) == target || getValue(matrix, right) == target {
		return true
	}
	var result bool

	for {
		mid := (left + right) / 2
		value := getValue(matrix, mid)
		//当左右边界无法在算出中间节点时，说明该数据不在区间中
		if mid == left || mid == right {
			result = false
			break
		}

		//中间值大于目标值，区间选择又区间,小于选择左区间,等于相当于找到了
		if value > target {
			right = mid
		} else if value < target {
			left = mid
		} else {
			result = true
			break
		}

	}

	return result
}

func getValue(matrix [][]int, index int) int {
	index = index
	n := len(matrix[0])

	x := index / n
	y := index % n

	return matrix[x][y]
}
