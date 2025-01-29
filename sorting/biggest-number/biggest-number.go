package biggestnumber

import (
	"fmt"
	"math"
	"sort"
)

type NumsAndDigit struct {
	num int
	pow float64
}

func solution(numbers []int) string {
	result := ""
	elements := [10][]NumsAndDigit{}
	for _, n := range numbers {
		var pow float64
		switch {
		case n < 10:
			pow = 0
		case n < 100:
			pow = 1
		case n < 1000:
			pow = 2
		default:
			pow = 3
		}
		idx := int(float64(n) / math.Pow(float64(10), pow))
		elements[idx] = append(elements[idx], NumsAndDigit{n, pow})
	}

	for _, ele := range elements {
		sort.Slice(ele, func(i, j int) bool {
			if ele[i].pow == ele[j].pow {
				return ele[i].num > ele[j].num
			}
			return ele[i].pow < ele[j].pow
		})
	}

	for i := len(elements) - 1; i >= 0; i-- {
		for _, ele := range elements[i] {
			result += fmt.Sprintf("%d", ele.num)
		}
	}

	return result
}

// 모범풀이

// func solution(numbers []int) string {
//     sort.Slice(numbers, func(i, j int) bool {
//         s1 := strconv.Itoa(numbers[i])
//         s2 := strconv.Itoa(numbers[j])
//         return s1+s2 > s2+s1
//     })

//     if numbers[0] == 0 {
//         return "0"
//     }
//     answer := ""
//     for i := range numbers {
//         answer += strconv.Itoa(numbers[i])
//     }
//     return answer
// }
