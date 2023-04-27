package leetCode

import "sort"

// map的方法
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

// 排序后比较，耗时大
func containsDuplicateOtherOne(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 哈希表
//
// 用一个 Set 存储所有的元素，如果元素已经存在于 Set 中，说明有重复，则返回 true
// 使用 Go 的空结构体 struct{}{} 来实现了 Set 的功能，它并没有成员变量，使用空结构体可以最小化空间占用
func containsDuplicateOtherTwo(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

// 双指针
//
// 将数组进行排序，然后使用双指针扫描，如果发现重复的元素，返回 true。
// 相比于哈希表，双指针解法需要先将数组排序，但是在实际应用中可能更快，因为哈希表需要消耗额外的空间。
func containsDuplicateOtherThree(nums []int) bool {
	sort.Ints(nums)
	left, right := 0, 1
	for right < len(nums) {
		if nums[left] == nums[right] {
			return true
		}
		left++
		right++
	}
	return false
}
