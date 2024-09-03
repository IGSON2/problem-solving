package network

func solution(n int, computers [][]int) int {
	cnt := 0
	dags := make([][]int, n)
	for i := 0; i < len(computers); i++ {
		for j := 0; j < len(computers); j++ {
			if i != j && computers[i][j] == 1 {
				dags[i] = append(dags[i], j)
			}
		}
	}
	checked := make(map[int]bool)

	for i := 0; i < n; i++ {
		cnt += dfs(checked, dags, i)
	}
	return cnt
}

func dfs(checked map[int]bool, dags [][]int, com int) int {
	if checked[com] {
		return 0
	}

	checked[com] = true
	for _, i := range dags[com] {
		dfs(checked, dags, i)
	}
	return 1
}
