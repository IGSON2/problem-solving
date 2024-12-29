package ps03

import (
	"strconv"
	"strings"
)

const (
	MO = iota
	TU
	WE
	TH
	FR
)

type Class struct {
	Day   int
	Start int
	End   int
}

func solution(schedule [][]string) int {
	// timeTable := [5][751]bool{}
	stack := make([]Class, 0)

	return dfs(schedule, stack)
}

func dfs(schedule [][]string, stack []Class) int {
	for subj := 0; subj < len(schedule); subj++ {

		for i := 0; i < len(schedule[0]); i++ {

			for len(stack) > 0 {
				v := stack[len(stack)-1]
				for _, next := range schedule[subj+1]{
					f convertSchedule(next)
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, convertSchedule(schedule[subj][i])...)
		}
	}
	return cnt
}

func convertSchedule(s string) []Class {
	result := make([]Class, 0)
	splited := strings.Split(s, " ")
	result = append(result, Class{convertDay(splited[0]), convertTime(splited[1]), 0})
	result[0].End = result[0].Start + 180
	if len(splited) > 2 {
		result = append(result, Class{convertDay(splited[2]), convertTime(splited[3]), 0})
		result[0].End, result[1].End = result[0].Start+90, result[1].Start+90
	}
	return result
}

func convertDay(s string) int {
	switch s {
	case "MO":
		return MO
	case "TU":
		return TU
	case "WE":
		return WE
	case "TH":
		return TH
	case "FR":
		return FR
	}
	return -1
}

func convertTime(s string) int {
	splited := strings.Split(s, ":")
	H, _ := strconv.Atoi(splited[0])
	H -= 9
	M, _ := strconv.Atoi(splited[1])
	return 60*H + M
}
