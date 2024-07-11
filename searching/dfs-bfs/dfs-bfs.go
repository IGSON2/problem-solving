package dfsbfs

import (
	"fmt"
	"problem-solving/utils"
	"sort"
)

// 자꾸 실패하나, 반례를 모르겠음
func dfsbfs() {
	f := utils.OpenFile("./testcase")
	defer f.Close()

	var N, M, V int
	fmt.Fscanf(f, "%d %d %d", &N, &M, &V)

	adjacencies := make(map[int][]int)
	for i := 0; i < M; i++ {
		var n, m int
		fmt.Fscanf(f, "%d %d\n", &n, &m)
		adjacencies[n] = append(adjacencies[n], m)
		adjacencies[m] = append(adjacencies[m], n)
	}

	for _, v := range adjacencies {
		sort.Ints(v)
	}

	// dfs
	dfs := make([]int, 0)
	stack := make([]int, 0)
	stack = append(stack, V)

	visited := make(map[int]bool)

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[v] {
			continue
		}

		visited[v] = true
		dfs = append(dfs, v)

		for i := len(adjacencies[v]) - 1; i >= 0; i-- {
			stack = append(stack, adjacencies[v][i])
		}

	}

	bfs := make([]int, 0)
	queue := make([]int, 0)
	queue = append(queue, V)
	visited = make(map[int]bool)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if visited[v] {
			continue
		}

		visited[v] = true
		bfs = append(bfs, v)

		for i := 0; i < len(adjacencies[v]); i++ {
			queue = append(queue, adjacencies[v][i])
		}

	}

	for i := 0; i < len(dfs); i++ {
		fmt.Print(dfs[i])
		if i == len(dfs)-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}

	for i := 0; i < len(bfs); i++ {
		fmt.Print(bfs[i])
		if i == len(bfs)-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}
