package journeytothemoon

import "testing"

func TestJourneyToTheMoon(t *testing.T) {
	tc := [][]int32{
		{1, 2},
		{3, 4},
	}

	got := journeyToMoon(1000000, tc)
	t.Log(got)
}
