package fibonaccinumber

func solution(n int) int {
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = (memo[i-2] + memo[i-1]) % 1234567
	}
	return memo[n]
}

// 직접 fibonacci 함수를 구현하기 보다, 메모라이제이션 하며 나아가는 게 훨씬 빠름
