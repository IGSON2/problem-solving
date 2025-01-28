package bestalbum

import "testing"

func TestSolution(t *testing.T) {
	genres := []string{"classic", "pop", "classic", "classic", "pop"}
	plays := []int{500, 600, 150, 800, 2500}
	result := solution(genres, plays)
	t.Log(result)
}

func TestMapExist(t *testing.T) {
	checked := make(map[string]int)
	t.Log(checked["hello"])
	// v,exist := checked["hello"]
}
