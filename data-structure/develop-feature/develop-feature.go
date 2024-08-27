package developfeature

import "math"

func solution(progresses []int, speeds []int) (result []int) {
	remainQ := make([]int, len(progresses))
	for i := 0; i < len(progresses); i++ {
		remainQ[i] = int(math.Ceil(float64(100-progresses[i]) / float64(speeds[i])))
	}
	for len(remainQ) > 0 {
		cnt := 1
		pop := remainQ[0]
		remainQ = remainQ[1:]

		for _, r := range remainQ {
			if pop >= r {
				remainQ = remainQ[1:]
				cnt++
			} else {
				break
			}
		}
		result = append(result, cnt)
	}
	return
}
