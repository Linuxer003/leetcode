package answer

func wordBreak(s string, wordDict []string) bool {
	// Initialization
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 这里 如果dp[j]为真 那么说明前面的字符串 可以分割为单词，不确定为几个
			//仅仅是s[0:j]可以正确分割
			//当s[0:j]可以分割，如果s[j:i]可以分割 那么 s[0:i] 也可以正确分割
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
