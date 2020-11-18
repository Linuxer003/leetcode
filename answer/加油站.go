package answer

func canCompleteCircuit(gas []int, cost []int) int {
	n, i := len(gas), 0
	for i < n {
		sumOfGas, sumOfCost := 0, 0
		// cnt表示从i从开始遍历n个。
		//如果遍历中间出现汽油和小于油耗，则跳出循环；如果遍历完n个，则说明可以绕环路行驶一周。
		cnt := 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfGas < sumOfCost {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			// 如果没有遍历完n个就结束了这轮，下一次就从i+cnt+1个加油站开始。
			i = i + cnt + 1
		}
	}
	return -1
}
