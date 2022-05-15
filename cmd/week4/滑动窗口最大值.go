package main

import (
	"container/heap"
	"fmt"
)

/**
239. 滑动窗口最大值

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
**/

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

//用二叉堆来实现
func maxSlidingWindow(nums []int, k int) []int {
	h := &list{
		nums: nums,
		set:  make([]int, 0, k),
	}
	result := make([]int, 0, len(nums)-k+1)

	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, i)
	}
	result = append(result, nums[h.GetTop()])

	for i := k; i < len(nums); i++ {
		heap.Push(h, i)

		//如果堆顶元素不在窗口内，全部取出(用下标信息来判断)
		for h.GetTop() <= i-k {
			heap.Pop(h)
		}

		//获取大根堆堆顶元素
		result = append(result, nums[h.GetTop()])
	}

	return result

}

type list struct {
	//存储所有数据
	nums []int
	//存储nums中的下标
	set []int
}

func (h *list) Len() int {
	return len(h.set)
}

func (h *list) Less(i, j int) bool {
	return h.nums[h.set[i]] > h.nums[h.set[j]]
}

func (h *list) Swap(i, j int) {
	h.set[i], h.set[j] = h.set[j], h.set[i]
}

func (h *list) Push(x interface{}) {
	h.set = append(h.set, x.(int))
}

func (h *list) Pop() interface{} {
	v := h.set[len(h.set)-1]
	h.set = h.set[:len(h.set)-1]

	return v
}

func (h *list) GetTop() int {
	return h.set[0]
}
