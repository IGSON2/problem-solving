package jadencase

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	testTxt := " 1AAA   2AAA 3AAA  4AAA 5AAA 6AAA  7AAA 8AAA 9AAA 0AAA"
	fmt.Printf("[%s]", solution(testTxt))
}
