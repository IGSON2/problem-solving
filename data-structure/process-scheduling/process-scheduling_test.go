package processscheduling

import (
	"fmt"
	"problem-solving/utils"
	"testing"
)

func testGeneratePriority() []int {
	s := make([]int, 5)
	for i := 0; i < len(s); i++ {
		s[i] = utils.MakeRanInt(0, 9)
	}
	return s
}

func TestProcess(t *testing.T) {
	priority := testGeneratePriority()
	location := utils.MakeRanInt(0, 5)
	fmt.Printf("priority : %v\nlocation : %d\n", priority, location)
	fmt.Println(solution(priority, location))
}
