package schedulinghomework

import (
	"strconv"
	"strings"
)

type schedule struct {
	subject string
	left    int
}

var scheduleMap = make(map[int]*schedule)

func convertTime(tStr string) int {
	tArr := strings.Split(tStr, ":")
	hh, _ := strconv.Atoi(tArr[0])
	mm, _ := strconv.Atoi(tArr[1])
	return hh*60 + mm
}

func SchedulingHomework(plans [][]string) []string {
	result := make([]string, 0)
	min := 1440

	for _, p := range plans {
		left, _ := strconv.Atoi(p[2])
		start := convertTime(p[1])
		if start < min {
			min = start
		}

		scheduleMap[start] = &schedule{
			subject: p[0],
			left:    left,
		}
	}

	stack := make([]*schedule, 0)
	processing := scheduleMap[min]
	for i := min + 1; len(result) != len(scheduleMap); i++ {
		processing.left--

		if v, ok := scheduleMap[i]; ok {
			if processing.left > 0 {
				stack = append(stack, processing)
			} else if processing.left == 0 {
				result = append(result, processing.subject)
			}
			processing = v
		} else {
			if processing.left == 0 {
				result = append(result, processing.subject)
				if len(stack) > 0 {
					processing = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return result
}
