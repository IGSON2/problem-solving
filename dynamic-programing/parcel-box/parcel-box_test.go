package parcelbox

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testCase := []int{1, 5, 3, 4, 2}

	result := solution(testCase)
	t.Log(result)
}
