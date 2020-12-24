package main

import (
	"fmt"
)

func main() {
	fmt.Println(magicalString(7))
}

func magicalString(n int) int {
	fast, ans := 2, 1
	if n == 0 {
		return 0
	}
	s := make([]byte, 0, n*2)
	s = append(s, '1', '2', '2')
	for i := 2; i < n; i++ {
		if s[i] == '2' {
			if s[fast] == '2' {
				s = append(s, '1', '1')
			} else {
				s = append(s, '2', '2')
			}
			fast += 2
		} else {
			if s[fast] == '2' {
				s = append(s, '1')
			} else {
				s = append(s, '2')
			}
			fast += 1
			ans++
		}
	}
	return ans
}
