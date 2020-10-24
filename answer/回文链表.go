package answer

func isPalindrome(head *ListNode) bool {
	var t []int
	for head != nil {
		t = append(t, head.Val)
		head = head.Next
	}

	for i, j := 0, len(t)-1; i < j; {
		if t[i] != t[j] {
			return false
		}
		i++
		j--
	}

	return true
}
