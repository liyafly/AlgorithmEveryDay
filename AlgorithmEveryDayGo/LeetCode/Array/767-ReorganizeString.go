package leetCode

// ReorganizeString
// https://leetcode.com/problems/reorganize-string/

func reorganizeString(S string) string {
	// 统计每个字符出现的次数
	count := [26]int{}
	for i := 0; i < len(S); i++ {
		count[S[i]-'a']++
	}
	max, letter := 0, 0
	// 找到出现次数最多的字符
	for i, c := range count {
		if c > max {
			max = c
			letter = i
		}
	}
	// 特殊情况 出现次数最多的字符数量大于了长度的一半
	if max > (len(S)+1)/2 {
		return ""
	}
	// 重新构造字符串
	res := make([]byte, len(S))
	idx := 0
	// 将出现次数最多的字符放到偶数位
	for count[letter] > 0 {
		res[idx] = byte(letter) + 'a'
		idx += 2
		count[letter]--
	}
	// 将其它字符放到奇数位
	for i, c := range count {
		for c > 0 {
			// 如果 idx 超过了字符串长度，从 1 开始
			if idx >= len(S) {
				idx = 1
			}

			res[idx] = byte(i) + 'a'
			idx += 2
			c--
		}
	}
	return string(res)
}
