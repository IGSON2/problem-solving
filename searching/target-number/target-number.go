package targetnumber

import "fmt"

func solution(numbers []int, target int) int {
	result := 0
	nums := dfs(0, 0, target, numbers)
	fmt.Println(nums)
	for _, num := range nums {
		if num == target {
			result++
		}
	}
	return result
}

func dfs(num, nextIdx, target int, numbers []int) []int {
	nums := make([]int, 0)
	if nextIdx >= len(numbers) {
		return []int{num}
	}

	pos := num + numbers[nextIdx]
	nums = append(nums, dfs(pos, nextIdx+1, target, numbers)...)

	neg := num - numbers[nextIdx]
	nums = append(nums, dfs(neg, nextIdx+1, target, numbers)...)
	return nums
}
