package main

import (
	"fmt"
	"leetcode/answer"
)

func main() {
	dis := [][]int{{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8}}
	fmt.Println(answer.PossibleBipartition(10, dis))
}
