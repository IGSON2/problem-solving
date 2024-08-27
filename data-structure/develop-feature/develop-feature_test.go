package developfeature

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	progresses := []int{93, 30, 55}
	speed := []int{1, 30, 5}
	fmt.Println(solution(progresses, speed))
}
