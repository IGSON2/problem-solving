package salesevent

func solution(want []string, number []int, discount []string) int {
	days := 0
	wishList := make(map[string]int)
	for i := 0; i < len(want); i++ {
		wishList[want[i]] = number[i]
	}

	for i := 0; i < 10; i++ {
		wishList[discount[i]]--
	}

	flag := true
	for _, v := range wishList {
		if v > 0 {
			flag = false
		}
	}
	if flag {
		days++
	}

Outer:
	for i := 10; i < len(discount); i++ {
		wishList[discount[i-10]]++
		wishList[discount[i]]--

		for _, v := range wishList {
			if v > 0 {
				continue Outer
			}
		}
		days++
	}
	return days
}

// b = 7,8,12
// a = 3,5,10
// r = 1,9
