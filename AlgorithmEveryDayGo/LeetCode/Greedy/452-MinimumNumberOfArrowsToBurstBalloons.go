package leetCode

import "sort"

// FindMinArrowShots
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(points [][]int) int {
	// 贪心算法
	// 按照气球的右边界排序，然后从左到右遍历气球
	// 如果当前气球的左边界大于上一支箭的位置，那么需要一支新的箭
	if len(points) == 0 {
		return 0
	}
	// 按照气球的右边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	cnt := 1
	// 第一支箭的位置
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		// 有重叠的，不需要新的箭
		if points[i][0] <= end {
			// 更新终点
			end = min(end, points[i][1])
		} else {
			// 没有重叠的，需要在新的位置放箭
			cnt++
			end = points[i][1]
		}
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
