package main

import "fmt"

//给数组代表的整数加一进位

func main() {
	fmt.Println(plusOne([]int{9, 9, 9, 8}))
}

func plusOne(digits []int) []int {
	a := 1
	k := len(digits) - 1

	for i := k; i >= 0; i-- {
		if digits[i]+a == 10 {
			digits[i] = 0
			a = 1
		} else {
			digits[i] = digits[i] + 1
			break
		}

		k--
	}

	if k < 0 {
		p := make([]int, 0, len(digits)+1)
		p = append(p, 1)

		for i := 0; i < len(digits); i++ {
			p = append(p, 0)
		}

		digits = p
	}

	return digits
}
