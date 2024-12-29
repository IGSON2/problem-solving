package leastcommonmultiple

import "sort"

func solution(arr []int) int {
	sort.Ints(arr)
	max := arr[len(arr)-1]
	for i := 1; ; i++ {
		LCM := max * i
		flag := true
		for _, a := range arr {
			if LCM%a != 0 {
				flag = false
			}
		}
		if flag {
			return LCM
		}
	}
}
