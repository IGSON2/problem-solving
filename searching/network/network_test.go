package network

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	coms := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(solution(len(coms), coms))
}
