package bestalbum

import "testing"

func TestSolution(t *testing.T) {
	genres := []string{"kpop", "kpop", "pop", "classic", "classic", "pop"}
	plays := []int{500, 500, 600, 150, 1000, 300}
	result := solution(genres, plays)
	t.Log(result)
}

func TestMapExist(t *testing.T) {
	checked := make(map[string]int)
	t.Log(checked["hello"])
	// v,exist := checked["hello"]
}
