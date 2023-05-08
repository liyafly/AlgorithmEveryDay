package leetCode

import "sort"

// FindContentChildren
// https://leetcode.com/problems/assign-cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		// 如果当前饼干满足当前小孩，那么 i 和 j 都加 1
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}
