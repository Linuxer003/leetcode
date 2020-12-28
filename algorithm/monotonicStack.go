package algorithm

func monotonicStack(heights []int) int {
	m := len(heights)
	if m == 0 {
		return 0
	}
	left, right := make([]int, m, m), make([]int, m, m)
	stack := make([]int, 0, m)
	for i := 0; i < m; i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0, m)
	for i := m - 1; i >= 0; i-- {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = m
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	temp := 0
	for i := 0; i < m; i++ {
		temp = (right[i] - left[i] - 1) * heights[i]
		if temp > ans {
			ans = temp
		}
	}
	return ans
}
