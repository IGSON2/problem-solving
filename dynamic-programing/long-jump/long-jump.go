package longjump

import (
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func combination(n, r int) *big.Int {
	if r > n {
		return big.NewInt(0)
	}
	num := factorial(n)
	denom := factorial(r)
	denom.Mul(denom, factorial(n-r))
	return num.Div(num, denom)
}

func solution(n int) int64 {
	var cnt int64 = 1
	for i := 1; i <= n-i; i++ {
		cnt += new(big.Int).Rem(combination(n-i, i), big.NewInt(1234567)).Int64()
	}
	return cnt
}
