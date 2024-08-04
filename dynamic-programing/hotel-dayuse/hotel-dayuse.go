package hoteldayuse

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Rent struct {
	start, end int
}

func solution(book_time [][]string) int {
	rents := make([]Rent, 0)
	for _, bt := range book_time {
		r := Rent{convertTime(bt[0]), convertTime(bt[1]) + 9}
		rents = append(rents, r)
	}

	timeTable := [1440]int{}

	for _, r := range rents {
		for i := r.start; i <= r.end && i < 1400; i++ {
			timeTable[i]++
		}
	}

	sort.Ints(timeTable[:])
	return timeTable[len(timeTable)-1]
}

func convertTime(timeStr string) int {
	splited := strings.Split(timeStr, ":")
	if len(splited) != 2 {
		panic(errors.New("invalid time string"))
	}
	hour, _ := strconv.Atoi(splited[0])
	minute, _ := strconv.Atoi(splited[1])
	return hour*60 + minute
}
