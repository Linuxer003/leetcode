package main

import (
	"fmt"
	"leetcode/answer"
)

func main() {
	head := &answer.ListNode{1, &answer.ListNode{2, &answer.ListNode{3, &answer.ListNode{4, &answer.ListNode{5, nil}}}}}
	head = answer.OddEvenList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
