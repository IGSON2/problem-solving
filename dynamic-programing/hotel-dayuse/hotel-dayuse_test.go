package hoteldayuse

import "testing"

func TestHotelDayuse(t *testing.T) {
	t.Log(solution([][]string{{"15:00", "17:00"}, {"16:40", "18:20"}, {"14:20", "15:20"}, {"14:10", "19:20"}, {"18:20", "21:20"}}))
}
