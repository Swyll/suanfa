package main

import "fmt"

/**
560. 和为 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
**/
func main() {
	fmt.Println(subarraySum2([]int{0, 1, 0}, 0))
}

//枚举
func subarraySum(nums []int, k int) int {
	var count, num int
	numLen := len(nums)

	for i := 0; i < numLen; i++ {
		num = 0

		for j := i; j < numLen; j++ {
			num = num + nums[j]
			if num == k {
				count++
			}
		}

	}

	return count
}

//前缀和+hash表(当前前缀和减去已有前缀和等于k说明有连续子数组和为k)prenum[i]=nums[0]+nums[1]+...+nums[i],prenum[i]-prenum[j]=k(j<i,所以prenum[i]-prenum[j]，代表的是j到i之间的子数组)
//顾推导prenum[i]-k存在则说明这个j到i之间的子数组和为k，这里用两个前缀数组之差表示一个连续数组，思想很nice
func subarraySum2(nums []int, k int) int {
	var num, count int
	cacheM := map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		num = num + nums[i]

		if v, ok := cacheM[num-k]; ok {
			count = count + v
		}

		cacheM[num] += 1
	}

	return count
}
