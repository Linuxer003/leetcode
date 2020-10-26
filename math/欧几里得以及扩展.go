package math

// Gcd 递归版本
func Gcd(a, b int) int {
	if a < b {
		return Gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Gcd2 迭代版本
func Gcd2(a, b int) int {
	if a < b {
		return Gcd2(b, a)
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		t := a
		a = b
		b = t % b
	}
	return a
}

// ExGcd 扩展欧几里得算法
// 最后结果的y就是b的逆元
func ExGcd(a, b int, x, y *int) int {
	if a < b {
		return ExGcd(b, a, x, y)
	}

	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	r := ExGcd(b, a%b, x, y)

	//x1 = y2, y1 = x2 - (a/b)*y2
	t := *y
	*y = *x - (a/b)**y
	*x = t
	return r
}

// ExGcd2 扩展欧几里得非递归版本
func ExGcd2(a, b int, x, y *int) int {
	if a < b {
		ExGcd2(b, a, x, y)
	}
	m, n := 0, 1
	*x, *y = 1, 0
	for b != 0 {
		var d, t = a / b, 0

		t = m
		m = *x - m*d
		*x = t

		t = n
		n = *y - n*d
		*y = t

		t = a % b
		a = b
		b = t
	}
	return a
}
