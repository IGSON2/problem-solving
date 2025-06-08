package connectingisland

import (
	"sort"
)

func solution(n int, costs [][]int) int {
	result := 0
	nodes := make(map[int][]int, 0)

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	for _, c := range costs {
		from, to, cost := c[0], c[1], c[2]
		if !isTraversal(nodes, from, to, make(map[int]bool)) {
			nodes[from] = append(nodes[from], to)
			nodes[to] = append(nodes[to], from)
			result += cost
		}
	}

	return result
}

func isTraversal(nodes map[int][]int, from, to int, visit map[int]bool) bool {
	if visit[from] {
		return false
	}

	visit[from] = true

	for _, v := range nodes[from] {
		if visit[v] {
			continue
		}

		if v == to {
			return true
		}

		if isTraversal(nodes, v, to, visit) {
			return true
		}
	}
	return false
}
