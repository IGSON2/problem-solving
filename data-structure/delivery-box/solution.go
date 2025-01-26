package deliverybox

func solution(order []int) int {
	loaded := 0
	container := make([]int, 0)
	subContainer := make([]int, 0)

	for i := 1; i < len(order); i++ {
		container = append(container, i)
	}

	for _, ord := range order {
		if len(subContainer) > 0 && subContainer[len(subContainer)-1] == ord {
			loaded++
			subContainer = subContainer[:len(subContainer)-1]
			continue
		}
		for i := 0; i < len(container); i++ {
			if container[i] == ord {
				loaded++
				container = container[i+1:]
				break
			}
			subContainer = append(subContainer, container[i])
			container = container[i+1:]
		}
	}
	return loaded
}
