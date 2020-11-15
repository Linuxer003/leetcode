package answer

import "strings"

func removeKDigits(num string, k int) string {
	stack := []byte(nil)
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// 此时已经删除了部分数字，k随着删除过程减小，但k未必减到0，所以还需要再删除一些。
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
