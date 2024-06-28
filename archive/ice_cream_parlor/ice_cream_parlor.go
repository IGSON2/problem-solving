package icecreamparlor

import "sort"

func icecreamParlor(m int32, arr []int32) []int32 {
	answer := make([]int32, 0)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
Outer:
	for i := 0; i < len(arr)-1; i++ {
		if m <= arr[i] {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if m-arr[i]-arr[j] == 0 {
				answer = append(answer, []int32{int32(i + 1), int32(j + 1)}...)
				break Outer
			}
		}
	}
	return answer
}
