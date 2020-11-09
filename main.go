package main

import (
	"fmt"
	"leetcode/answer"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	prices := [][]int{{1, 3}, {-2, 2}, {5, -1}, {-2, 4}}
	fmt.Println(answer.KClosest(prices, 2))
	fmt.Println(time.Now().UnixNano() - t)
}
