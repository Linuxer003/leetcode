package answer

func backspaceCompare(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	// 为了把下标 一正 一负的情况扔进循环里，所以用逻辑与
	for i >= 0 || j >= 0 {
		for i >= 0 {
			// 寻找第一个可以比较的下标
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
			// 出现一正一负的情况  直接返回false
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
