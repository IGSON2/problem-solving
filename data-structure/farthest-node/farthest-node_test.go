package farthestnode

import "testing"

func TestSolution(t *testing.T) {
	tc := [][]int{{3, 6}, {4, 3}, {3, 2}, {1, 3}, {1, 2}, {2, 4}, {5, 2}}
	// t.Log(solution(6, tc))
	t.Log(HeapSolution(6, tc))
}
