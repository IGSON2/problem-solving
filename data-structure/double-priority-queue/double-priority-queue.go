package doublepriorityqueue

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func solution(operations []string) []int {
	dequeue := make([]int, 0)
	for _, op := range operations {
		command := strings.Split(op, " ")
		switch command[0] {
		case "I":
			num, err := strconv.Atoi(command[1])
			if err != nil {
				log.Panicln(err)
			}
			dequeue = append(dequeue, num)
			sort.Slice(dequeue, func(i, j int) bool {
				return dequeue[i] < dequeue[j]
			})
		case "D":
			if len(dequeue) == 0 {
				continue
			}
			switch command[1] {
			case "1":
				dequeue = dequeue[:len(dequeue)-1]
			case "-1":
				dequeue = dequeue[1:]
			}
		}
	}

	if len(dequeue) == 0 {
		dequeue = []int{0}
	}

	return []int{dequeue[len(dequeue)-1], dequeue[0]}
}
