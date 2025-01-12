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
