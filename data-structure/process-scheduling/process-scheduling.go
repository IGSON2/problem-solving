package processscheduling

type process struct {
	index, value int
}

func solution(priorities []int, location int) int {
	cnt := 1

	pq := make([]process, len(priorities))
	for i, v := range priorities {
		pq[i] = process{i, v}
	}

Outer:
	for len(pq) > 0 {
		pop := pq[0]
		pq = pq[1:]
		for _, v := range pq {
			if pop.value < v.value {
				pq = append(pq, pop)
				continue Outer
			}
		}
		if pop.index == location {
			return cnt
		}
		cnt++
	}
	panic("error")
}
