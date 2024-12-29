package hindex

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	citation := []int{1, 2, 3, 4, 6, 7, 8, 8, 8, 11, 15, 28, 36, 9}
	fmt.Println(solution(citation))
}
