package main

import (
	"sort"
)

/**
49. 字母异位词分组

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
**/

func main() {
	a := "ssddd"
	sort.Sort(String(a))
}

type String []rune

func (s String) Len() int           { return len(s) }
func (s String) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s String) Less(i, j int) bool { return s[i] < s[j] }

func groupAnagrams(strs []string) [][]string {
	cacheM := make(map[string][]string)
	cache := make([][]string, 0, 4)

	for _, v := range strs {
		b := []rune(v)
		sort.Sort(String(b))

		if ss, ok := cacheM[string(b)]; ok {
			cacheM[string(b)] = append(ss, v)
		} else {
			cacheM[string(b)] = []string{v}
		}
	}

	for _, v := range cacheM {
		cache = append(cache, v)
	}

	return cache
}
