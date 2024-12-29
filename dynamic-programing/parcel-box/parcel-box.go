package parcelbox

func solution(order []int) int {
	result := 1
	startIdx := order[0] - 1

	boxes := make([]int, len(order))
	for i := 0; i < len(boxes); i++ {
		boxes[i] = i + 1
	}

	checked, remain := boxes[:startIdx], boxes[startIdx+1:]

	for len(remain) > 0 || len(checked) > 0 {
		curOrder := order[result]
		if len(checked) > 0 && curOrder == checked[len(checked)-1] {
			checked = checked[:len(checked)-1]
			result++
		} else if len(remain) > 0 && curOrder == remain[0] {
			remain = remain[1:]
			result++
		} else {
			if len(remain) > 0 {
				checked = append(checked, remain[0])
				remain = remain[1:]
				continue
			}
			break
		}
	}

	return result
}

// 더 간결한 풀이
func betterSolution(order []int) int {
	answer := 0
	list := make([]int, len(order))
	for i, v := range order { // order가 [1, 5, 3, 4, 2]면 list는 [0, 2, 3, 4, 1]
		list[v-1] = i
	}
	var s []int
	for _, v := range list {
		if v == answer {
			answer += 1
		} else {
			s = append(s, v)
		}
		for 0 < len(s) && s[len(s)-1] == answer {
			answer += 1
			s = s[:len(s)-1]
		}
	}
	return answer
}
