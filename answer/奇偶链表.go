package answer

func oddEvenList(head *ListNode) *ListNode {
	var odd, even, oddNext, evenNext *ListNode
	odd = oddNext
	even = evenNext
	i := 1
	for head != nil {
		if i%2 == 1 {
			oddNext = head
			oddNext = oddNext.Next
			head = head.Next
		} else {
			evenNext = head
			evenNext = evenNext.Next
			head = head.Next
		}
		i++
	}
	oddNext = even
	head = odd
	return head
}
