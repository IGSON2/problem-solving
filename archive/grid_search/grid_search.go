package gridsearch

import "strings"

func gridSearch(G []string, P []string) string {
	var answer string = "NO"
Outer:
	for i, g := range G {
		if len(G)-i < len(P) {
			break
		}
		if exist := strings.Index(g, P[0]); exist != -1 {
			curIdx := -2
		Inner:
			for k := 0; k < len(g)-len(P[0])+1; k++ {
				cnt := 0
				for j, p := range P {
					idx := strings.Index(G[i+j][k:], p) + k
					if idx-k == -1 {
						continue Inner
					}
					if j == 0 {
						curIdx = idx
						cnt++
						continue
					}
					if curIdx != idx {
						continue Inner
					}
					cnt++
				}
				if cnt == len(P) {
					answer = "YES"
					break Outer
				}
			}
		}
	}
	return answer
}
