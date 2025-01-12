package convertnumber

type xAndCnt struct {
	x   int
	cnt int
}

func solution(x int, y int, n int) int {
	queue := []xAndCnt{{x, 0}}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v.x == y {
			return v.cnt
		}

		for _, cal := range []int{v.x * 2, v.x * 3, v.x + n} {
			if cal <= y && !visited[cal] {
				visited[cal] = true
				next := xAndCnt{cal, v.cnt + 1}
				queue = append(queue, next)
			}
		}
	}

	return -1
}
