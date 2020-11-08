package answer

import "sort"

var Bit = [1e4 + 1]int{}

// init 采用位运算+递推的方式 计算二进制中1的个数
func Init() {
	for i := 1; i <= 1e4+1; i++ {
		//当前数字的含1数量 = 比目标少一位的含1数量 + 当前数字的最后一位与1与运算的结果
		Bit[i] = Bit[i>>1] + i&1
	}
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		cx, cy := Bit[x], Bit[y]
		return cx < cy || cx == cy && x < y
	})
	return arr
}
