package leetCode

// RemoveDuplicatesFromSortedArray
// https://leetcode.com/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
