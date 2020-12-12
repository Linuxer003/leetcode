package answer

func WiggleMaxLength(nums []int) int {
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	n, down, up := len(nums), 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = max(down+1, up)
		} else if nums[i] < nums[i-1] {
			down = max(up+1, up)
		}
	}
	return max(up, down)
}
