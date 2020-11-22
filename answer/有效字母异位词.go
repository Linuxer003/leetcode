package answer

import "sort"

func isAnagram1(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}
	m := make(map[byte]int, lenS)
	for i := 0; i < lenS; i++ {
		m[s[i]]++
	}
	for i := 0; i < lenT; i++ {
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	bytesS := []byte(s)
	bytesT := []byte(t)
	sort.Slice(bytesS, func(i, j int) bool {
		return bytesS[i] < bytesS[j]
	})
	sort.Slice(bytesT, func(i, j int) bool {
		return bytesT[i] < bytesT[j]
	})
	s = string(bytesS)
	t = string(bytesT)
	return s == t
}
