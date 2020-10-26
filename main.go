package main

import (
	"leetcode/answer"
)

func main() {
	var x, y int
	answer.ExGcd2(47, 30, &x, &y)
	println(y)
}
