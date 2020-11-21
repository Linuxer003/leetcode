package main

import (
	"fmt"
	"leetcode/answer"
)

func main() {
	head := &answer.ListNode{4, &answer.ListNode{2, &answer.ListNode{1, &answer.ListNode{3, nil}}}}
	answer.SortList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
