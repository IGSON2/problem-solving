package salesevent

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	want := []string{"banana", "apple", "rice", "pork", "pot"}
	number := []int{3, 2, 2, 2, 1}
	discount := []string{"chicken", "apple", "apple", "banana", "rice", "apple", "pork", "banana", "pork", "rice", "pot", "banana", "apple", "banana"}
	fmt.Printf("%d", solution(want, number, discount))
}
