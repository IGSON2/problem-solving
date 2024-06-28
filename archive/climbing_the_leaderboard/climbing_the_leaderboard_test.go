package climbingtheleaderboard

import "testing"

func TestClimbingTheLeaderboard(t *testing.T) {
	rankArray := []int32{100, 100, 50, 40, 40, 20, 10}
	playerArray := []int32{5, 25, 50, 120}
	result := climbingLeaderboard(rankArray, playerArray)
	t.Log(result)
}
