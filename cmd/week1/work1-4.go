package main

import "fmt"

//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
func main() {
	fmt.Println(maximalRectangle([][]byte{[]byte{'1', '1'}, []byte{'1', '1'}}))
}

//以二维数组每行最左边为原点，算出每个数组的连续长度，并由矩形最大面积算法获取最大面积
func maximalRectangle(matrix [][]byte) int {
	cacheM := make(map[int]int)
	var max int

	for y := 0; y < len(matrix); y++ {
		cacheS := make([]int, 0, len(matrix[y]))

		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == '1' {
				cacheM[x] += 1
			} else {
				cacheM[x] = 0
			}

			cacheS = append(cacheS, cacheM[x])
		}

		ca := largestRectangleArea(cacheS)
		if ca > max {
			max = ca
		}
	}

	return max
}

func largestRectangleArea(heights []int) int {
	length := len(heights)
	left, right := make([]int, length), make([]int, length)
	stack := []int{}

	for i := 0; i < length; i++ {
		//如果栈不为空，且栈顶元素小于当前值，出栈直到栈中小于当前值的元素全部出栈之后在入栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		//如果栈为空，那么说明左边没有值小于当前值，左边界就为数组最左边
		if len(stack) == 0 {
			left[i] = 0
		} else {
			left[i] = stack[len(stack)-1] + 1
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := length - 1; i >= 0; i-- {
		//如果栈不为空，且栈顶元素小于当前值，出栈直到栈中小于当前值的元素全部出栈之后在入栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		//如果栈为空，那么说明右边没有值大于当前值，右边界就为数组最右边
		if len(stack) == 0 {
			right[i] = length - 1
		} else {
			right[i] = stack[len(stack)-1] - 1
		}

		stack = append(stack, i)
	}

	var max int
	for i := 0; i < length; i++ {
		ca := (right[i] - left[i] + 1) * heights[i]
		if ca > max {
			max = ca
		}
	}

	return max
}
