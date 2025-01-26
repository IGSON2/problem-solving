package deliverybox

import "testing"

func TestSolution(t *testing.T) {
	tcs := [][]int{
		{4, 3, 1, 2, 5},
		{5, 4, 3, 2, 1},
	}
	for _, tc := range tcs {
		t.Log(solution(tc))
	}
}
