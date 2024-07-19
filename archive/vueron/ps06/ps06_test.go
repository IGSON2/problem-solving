package ps06

import "testing"

func TestSqrt(t *testing.T) {
	t.Log(sqrt(10))
}

func TestPs06(t *testing.T) {
	arr := [][2]int{{-1, -2}, {-1, 2}, {2, 2}, {3, -3}, {1, 1}}
	t.Log(ps06(arr))
}
