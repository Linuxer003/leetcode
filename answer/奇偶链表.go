package answer

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
