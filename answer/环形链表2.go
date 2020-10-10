package answer

/**
142
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				slow = slow.Next
				p = p.Next
			}
			return slow
		}
	}
	return nil
}
