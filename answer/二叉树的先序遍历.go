package answer

// PreorderTraversal 先序遍历的迭代实现
func PreorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ans []int
	if root == nil {
		return nil
	} else {
		// 这里因为要实现先序  所以出栈要先出左节点，所以先进右节点
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		ans = append(ans, root.Val)
	}
	for len(stack) != 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
		ans = append(ans, t.Val)
	}
	return ans
}
