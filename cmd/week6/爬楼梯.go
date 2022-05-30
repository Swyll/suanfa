package main

import "fmt"

/**
70.爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
**/
func main() {
	fmt.Println(climbStairs2(4))
}

func climbStairs(n int) int {
	var result int
	var fn func(int)
	op := []int{1, 2}

	fn = func(curr int) {
		if curr >= n {
			if curr == n {
				result++
			}

			return
		}

		//决策
		for _, v := range op {
			fn(curr + v)
		}

	}
	fn(0)

	return result
}

//金典动态规划题,设最后一级为f(x)=f(x-1)+f(x-2)由最后一阶和最后第二阶和解得
//相当于在状态节点相同的点合并两节点和
//前一个状态为f(x-1),f(x-2),决策为走1步还是2步
func climbStairs2(n int) int {
	//p为f(x-1),q为f(x-2)
	r, q, p := 1, 0, 1

	for i := 2; i <= n; i++ {
		q = p
		p = r
		r = p + q
	}

	return r
}
