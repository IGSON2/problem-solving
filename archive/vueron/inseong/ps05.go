package inseong

import "fmt"

type cell5 struct {
	i, j int
}

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isSafe(i, j int, N int) bool {
	return i >= 0 && i < N && j >= 0 && j < N
}

func Ps05(N int, arr [][]uint8) {
	adj := make(map[cell5][]cell5)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if isSafe(ni, nj, N) {
					adj[cell5{i, j}] = append(adj[cell5{i, j}], cell5{ni, nj})
				}
			}
		}
	}

	result := make([]int, 0)
	visited := make(map[cell5]bool)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			complex := dfs(arr, cell5{i, j}, adj, visited)
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

func dfs(arr [][]uint8, c cell5, adj map[cell5][]cell5, visited map[cell5]bool) int {
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
