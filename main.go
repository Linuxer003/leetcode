package main

import (
	"fmt"
	"leetcode/answer"
)

func main() {
	root := &answer.TreeNode{1, nil, &answer.TreeNode{2, &answer.TreeNode{3, nil, nil}, nil}}

	fmt.Println(answer.PreorderTraversal(root))
}
