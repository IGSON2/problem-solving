package continuousarray

func solution(elements []int) int {
	elements = append(elements, elements...)

	dup := make(map[int]struct{})

	for i := 1; i <= len(elements)/2; i++ {
		for j := 0; j < len(elements)/2; j++ {
			sum := 0
			for _, e := range elements[j : j+i] {
				sum += e
			}
			dup[sum] = struct{}{}
		}
	}

	return len(dup)
}
