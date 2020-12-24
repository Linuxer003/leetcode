package algorithm

const mod = 1e9 + 7

type SegmentTree struct {
	tree []int
	lazy []int
}

func NewSegmentTree(n int) SegmentTree {
	return SegmentTree{
		tree: make([]int, n*4),
		lazy: make([]int, n*4),
	}
}

func (t *SegmentTree) pushDown(pos, left, right int) {
	if t.lazy[pos] == 0 {
		return
	}
	t.tree[pos] += t.lazy[pos] * (right - left + 1)
	if left < right {
		t.lazy[pos*2+1] += t.lazy[pos]
		t.lazy[pos*2+2] += t.lazy[pos]
	}
	t.lazy[pos] = 0
}

func (t *SegmentTree) update(pos, left, right, curLeft, curRight, val int) {
	t.pushDown(pos, left, right)
	if right < curLeft || left > curRight {
		return
	}
	if curLeft <= left && right <= curRight {
		t.lazy[pos] = val
		t.pushDown(pos, left, right)
		return
	}
	mid := (left + right) >> 1
	t.update(pos*2+1, left, mid, curLeft, curRight, val)
	t.update(pos*2+2, mid+1, right, curLeft, curRight, val)
	t.tree[pos] = (t.tree[pos*2+1] + t.tree[pos*2+2]) % mod
}

func (t *SegmentTree) query(pos, left, right, curLeft, curRight int) int {
	if right < curLeft || left > curRight {
		return 0
	}
	t.pushDown(pos, left, right)
	if curLeft <= left && right <= curRight {
		return t.tree[pos]
	}
	mid := (left + right) >> 1
	return t.query(pos*2+1, left, mid, curLeft, curRight) +
		t.query(pos*2+2, mid+1, right, curLeft, curRight)
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	graph := map[int][]int{} // leader --> followers
	for _, l := range leadership {
		graph[l[0]] = append(graph[l[0]], l[1])
	}
	// 使用dfs将题目中的节点转换为线段树里的，题目里的数字只能算是"value"
	// 而线段树里的需要以map的key为题目中的数字，value为线段树的id
	L, R := map[int]int{}, map[int]int{}
	id := -1
	var dfs func(leader int)
	dfs = func(leader int) {
		id++
		L[leader] = id
		for _, follower := range graph[leader] {
			dfs(follower)
		}
		R[leader] = id
	}
	dfs(1)
	// 使用线段树执行操作
	t := NewSegmentTree(n)
	result := []int{}
	for _, op := range operations {
		switch op[0] {
		case 1:
			t.update(0, 0, n-1, L[op[1]], L[op[1]], op[2])
			break
		case 2:
			t.update(0, 0, n-1, L[op[1]], R[op[1]], op[2])
			break
		case 3:
			result = append(result, t.query(0, 0, n-1, L[op[1]], R[op[1]])%mod)
			break
		}
	}
	return result
}
