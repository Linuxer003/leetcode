package answer

/**
143
*/
func reorderList(head *ListNode) {
	reverse(head)
	nodeArray = nodeArray[:len(nodeArray)-1]
	if len(nodeArray) == 0 {
		return
	}
	p, q := head, &ListNode{}
	for p.Next != nodeArray[0] {
		q.Next = p.Next
		q.Next.Next = nodeArray[0]
		q = q.Next.Next
		p = p.Next
		nodeArray = nodeArray[:len(nodeArray)-1]
	}
	head = q
}

var nodeArray []*ListNode

func reverse(head *ListNode) {
	if head.Next != nil {
		reverse(head.Next)
	}
	nodeArray = append(nodeArray, head)
}
