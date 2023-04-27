package leetCode

// RotateArray
// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	// 取模，处理 k 大于数组长度的情况
	k %= len(nums)
	// 反转整个数组
	reverse(nums, 0, len(nums)-1)
	// 反转0 到k个
	reverse(nums, 0, k-1)
	// 反转后面的
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, begin, end int) {
	for begin < end {
		nums[begin], nums[end] = nums[end], nums[begin]
		begin++
		end--
	}
}

func rotateOther(num []int, k int) {
	length := len(num)
	k %= length
	if k == 0 || length == 1 {
		return
	}
	temp, count := num[0], 0 // 用 temp 存放需要移动的元素，用 count 表示当前已经处理了几个位置
	// i 从 k 开始，每次增加 k，模拟移动 k 步的效果
	for i, cnt := k, 0; cnt < length; i += k {
		t := num[i%length]
		num[i%length] = temp
		temp = t
		if i%length == count {
			count++   // 更新处理的元素个数
			i = count // 重新设置 i 的值，从下一个位置开始移动
			temp = num[i%length]
		}
		cnt++ // 更新处理的元素个数
	}
}
