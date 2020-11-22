package answer

import "sort"

func SortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var arr []*ListNode
	for head != nil {
		t := head
		arr = append(arr, t)
		head = head.Next
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	n := len(arr)

	for i := 0; i < n-1; i++ {
		arr[i].Next = arr[i+1]
	}
	// 这里必须执行末尾置空的操作，不然会出现无穷无尽的尾巴
	arr[n-1].Next = nil
	return arr[0]
}
