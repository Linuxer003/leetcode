package answer

// SumNumbers
/**
初始时，将根节点和根节点的值分别加入两个队列。每次从两个队列分别取出一个节点和一个数字，进行如下操作：
	1.如果当前节点是叶子节点，则将该节点对应的数字加到数字之和；
	2.如果当前节点不是叶子节点，则获得当前节点的非空子节点，并根据当前节点对应的数字和子节点的值计算子节点对应的数字，然后将子节点和子节点对应的数字分别加入两个队列。
搜索结束后，即可得到所有叶子节点对应的数字之和。
*/
//func SumNumbers(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	node, num, ans := []*TreeNode{}, []int{}, 0
//	node = append(node, root)
//	num = append(num, root.Val)
//	for len(node) != 0 {
//		t := node[0]
//		node = node[1:]
//		n := num[0]
//		num = num[1:]
//		if t.Left == nil && t.Right == nil {
//			ans += n
//		}
//		if t.Left != nil {
//			node = append(node, t.Left)
//			num = append(num, n*10+t.Left.Val)
//		}
//		if t.Right != nil {
//			node = append(node, t.Right)
//			num = append(num, n*10+t.Right.Val)
//		}
//	}
//	return ans
//}

// sumNumbers
/**
递归解法
*/
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
