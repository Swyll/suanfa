package main

import (
	"strconv"
	"strings"
)

/**
网站域名 "discuss.leetcode.com" 由多个子域名组成。顶级域名为 "com" ，二级域名为 "leetcode.com" ，最低一级为 "discuss.leetcode.com" 。当访问域名 "discuss.leetcode.com" 时，同时也会隐式访问其父域名 "leetcode.com" 以及 "com" 。
计数配对域名 是遵循 "rep d1.d2.d3" 或 "rep d1.d2" 格式的一个域名表示，其中 rep 表示访问域名的次数，d1.d2.d3 为域名本身。
例如，"9001 discuss.leetcode.com" 就是一个 计数配对域名 ，表示 discuss.leetcode.com 被访问了 9001 次。
给你一个 计数配对域名 组成的数组 cpdomains ，解析得到输入中每个子域名对应的 计数配对域名 ，并以数组形式返回。可以按 任意顺序 返回答案。
**/
func main() {
}

func subdomainVisits(cpdomains []string) []string {
	cacheM := make(map[string]int)

	for _, v := range cpdomains {
		vs := strings.Split(v, " ")

		num, _ := strconv.Atoi(vs[0])

		vss := strings.Split(vs[1], ".")
		vssLen := len(vss)
		cache := ""

		for i := vssLen - 1; i >= 0; i-- {
			if i == vssLen-1 {
				cache = vss[i]
			} else {
				cache = vss[i] + "." + cache
			}

			if v3, ok := cacheM[cache]; ok {
				cacheM[cache] = v3 + num
			} else {
				cacheM[cache] = num
			}
		}
	}

	cacheSlice := make([]string, 0, len(cacheM))
	for k, v := range cacheM {
		str := strconv.Itoa(v) + " " + k
		cacheSlice = append(cacheSlice, str)
	}

	return cacheSlice
}
