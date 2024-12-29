package matchingprediction

import "math"

func solution(n int, a int, b int) int {
	answer := 0
	rounds := int(math.Log2(float64(n)))

	for i := 0; i <= rounds; i++ {
		if a == b {
			answer = i
			break
		}
		a = int(math.Ceil(float64(a) / float64(2)))
		b = int(math.Ceil(float64(b) / float64(2)))
	}
	return answer
}
