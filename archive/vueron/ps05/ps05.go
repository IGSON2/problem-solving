package ps05

import "fmt"

type cell struct {
	i, j int
}

var direction = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isSafe(i, j int, N int) bool {
	return i >= 0 && i < N && j >= 0 && j < N
}

func ps05(N int, arr [][]uint8) {
	adj := make(map[cell][]cell)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for _, d := range direction {
				ni, nj := i+d[0], j+d[1]
				if isSafe(ni, nj, N) {
					adj[cell{i, j}] = append(adj[cell{i, j}], cell{ni, nj})
				}
			}
		}
	}

	result := make([]int, 0)
	visited := make(map[cell]bool)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			complex := dfs(arr, cell{i, j}, adj, visited)
			if complex > 0 {
				result = append(result, complex)
			}
		}
	}

	sort(result)

	fmt.Println(len(result))
	for _, r := range result {
		fmt.Println(r)
	}
}

func dfs(arr [][]uint8, c cell, adj map[cell][]cell, visited map[cell]bool) int {
	if visited[c] || arr[c.i][c.j] == 0 {
		return 0
	}

	cnt := 1
	visited[c] = true
	for _, n := range adj[c] {
		if !visited[n] {
			cnt += dfs(arr, n, adj, visited)
		}
	}
	return cnt
}

func sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
