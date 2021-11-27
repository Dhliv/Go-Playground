package Tests

func Sum(a, b int) int {
	return a + b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibo(n int) int64 {
	if n == 1 || n == 0 {
		return int64(1)
	}
	return Fibo(n-1) + Fibo(n-2)
}
