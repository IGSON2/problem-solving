package routingtrip

import "testing"

func TestRoutingTrip(t *testing.T) {
	// tickets := [][]string{{"ICN", "JFK"}, {"HND", "IAD"}, {"JFK", "HND"}}
	tickets := [][]string{{"ICN", "AAA"}, {"ICN", "BBB"}, {"BBB", "ICN"}} // 반례 해겲 못함
	result := solution(tickets)
	t.Log(result)
}

// func TestSlice(t *testing.T) {
// 	s := []int{1}
// 	s = s[1:]
// 	t.Log(s) // []
// }
