package ricochetbot

import "testing"

func TestRicochet(t *testing.T) {
	board := []string{
		"...D..R",
		".D.G...",
		"....D.D",
		"D....D.",
		"..D....",
	}
	result := solution(board)
	t.Log(result)
}
