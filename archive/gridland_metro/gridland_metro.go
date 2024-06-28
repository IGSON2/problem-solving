package gridlandmetro

type trackRange struct {
	min int32
	max int32
}

func (t *trackRange) updatePoints(start, end int32) bool {
	if start >= t.max || end <= t.min {
		if start == t.max {
			t.max--
		}
		if end == t.min {
			t.min++
		}
		return false
	}

	if start < t.min {
		t.min = start
	}

	if end > t.max {
		t.max = end
	}
	return true
}

func gridlandMetro(n int32, m int32, k int32, track [][]int32) uint64 {
	wholeArea := uint64(n) * uint64(m)
	var trackArea uint64
	rangeByRow := make(map[int32][]*trackRange)
	for _, t := range track {
		var isInRange bool
		v, exist := rangeByRow[t[0]]
		if exist {
			for _, v2 := range v {
				isInRange = v2.updatePoints(t[1], t[2])
			}
		}

		if !isInRange {
			newRange := &trackRange{t[1], t[2]}
			rangeByRow[t[0]] = append(rangeByRow[t[0]], newRange)
		}
	}

	for _, v := range rangeByRow {
		for _, v2 := range v {
			trackArea += uint64(v2.max - v2.min + 1)
		}
	}
	return wholeArea - trackArea
}
