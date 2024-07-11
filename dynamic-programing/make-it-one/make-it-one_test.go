package makeitone

import "testing"

func TestMakeItOne(t *testing.T) {
	makeItOne(10)
}

// 120 -> 40 -> 39 -> 13 -> 12 -> 4 -> 2 -> 1
// 120 -> 40 -> 20 -> 10 -> 9 -> 3 -> 1
