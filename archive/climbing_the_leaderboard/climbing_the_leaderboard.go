package climbingtheleaderboard

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	discinctRanks := make([]int32, 0)
	var rank int32 = 1

	discinctRanks = append(discinctRanks, ranked[0])
	for _, rScore := range ranked[1:] {
		if latest := discinctRanks[len(discinctRanks)-1]; latest != rScore {
			discinctRanks = append(discinctRanks, rScore)
			rank++
		}
	}
	resultRanks := make([]int32, 0)

	for _, p := range player {
		var prev int32
		for i, r := range discinctRanks {
			if r > p {
				prev = int32(i + 1)
			} else {
				break
			}
		}
		resultRanks = append(resultRanks, prev+1)
	}

	return resultRanks
}
