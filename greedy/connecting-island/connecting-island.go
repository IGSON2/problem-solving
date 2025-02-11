package connectionisland

import "sort"

func solution(n int, costs [][]int) int {
	result := 0
	parents := make(map[int][]int)

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	for _, c := range costs {
		from, to, cost := c[0], c[1], c[2]
		if !isTraversal(parents, from, to, make(map[int]bool)) {
			result += cost
		}
	}

	return result
}

func isTraversal(parents map[int][]int, from, to int, visit map[int]bool) bool {
	if visit[from] {
		return false
	}

	routes, exist := parents[from]
	if !exist {
		return false
	}

	for _, r := range routes {
		if visit[r] {
			return false
		}

		if r == to {
			return true
		}

		if isTraversal(parents, r, to, visit) {
			return true
		}
	}
	return false
}
