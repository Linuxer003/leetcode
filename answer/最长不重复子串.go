package answer

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int(nil)
	n := len(s)
	// 右指针，初始值为-1， 相当于我们在字符串的左边界的左侧，还没开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第i到rk个字符是一个极长地无重复字符子串
		ans = max2(ans, rk-i+1)
	}
	return ans
}

func max2(x, y int) int {
	if x < y {
		return y
	}
	return x
}
