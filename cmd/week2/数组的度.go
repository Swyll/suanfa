package main

/**
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
**/

func main() {
}

type node struct {
	start int
	end   int
	count int
}

func findShortestSubArray(nums []int) int {
	cacheM := make(map[int]*node)

	for i, v := range nums {
		if v2, ok := cacheM[v]; ok {
			v2.end = i
			v2.count++
		} else {
			cacheM[v] = &node{
				start: i,
				end:   i,
				count: 1,
			}
		}
	}

	var max, max2 int

	for _, v := range cacheM {
		if v.count > max {
			max = v.count
			max2 = v.end - v.start + 1
		} else if v.count == max {
			if v.end-v.start+1 < max2 {
				max2 = v.end - v.start + 1
			}
		}
	}

	return max2
}
