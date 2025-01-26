package rollcake

func solution(topping []int) int {
	cnt := 0
	chulsoo := make(map[int]int)
	brother := make(map[int]int)

	for _, t := range topping {
		brother[t]++
	}

	for _, t := range topping {
		chulsoo[t]++
		brother[t]--
		if brother[t] == 0 {
			delete(brother, t)
		}
		if len(chulsoo) == len(brother) {
			cnt++
		}
	}
	return cnt
}
