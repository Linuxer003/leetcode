package main

import (
	"leetcode/math"
)

func main() {
	var x, y int
	math.ExGcd2(47, 30, &x, &y)
	println(y)
}
