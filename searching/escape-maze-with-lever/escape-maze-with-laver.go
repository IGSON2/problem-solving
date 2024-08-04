package escapemazewithlever

import (
	"strings"
)

type Cell struct {
	i, j int
}

func (c Cell) isSafe(N, M int) bool {
	return c.i >= 0 && c.j >= 0 && c.i < N && c.j < M
}

var direction = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solution(maps []string) int {
	var start, lever, end Cell
	var N, M = len(maps), len(maps[0])
	map2D := make([][]string, N)

	for i, line := range maps {
		splited := strings.Split(line, "")
		map2D[i] = splited
	}

	adjacencies := make(map[Cell][]Cell)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			c := Cell{i, j}
			switch string(maps[i][j]) {
			case "S":
				start = c
			case "L":
				lever = c
			case "E":
				end = c
			}

			for _, d := range direction {
				adj := Cell{c.i + d[0], c.j + d[1]}
				if adj.isSafe(N, M) {
					adjacencies[c] = append(adjacencies[c], adj)
				}
			}
		}
	}
	visited := make(map[Cell]bool)
	toLever := bfs(maps, start, lever, adjacencies, visited)

	visited = make(map[Cell]bool)
	toEnd := bfs(maps, lever, end, adjacencies, visited)

	if toLever == 0 || toEnd == 0 {
		return -1
	}

	return toLever + toEnd
}

// 최단거리 구하는 문제는 bfs로 풀어야함. dfs의 첫 탐색 결과가 최단거리가 아닐 수 있기 때문
func bfs(maps []string, start, end Cell, adjacencies map[Cell][]Cell, visited map[Cell]bool) int {
	distance := make(map[Cell]int)
	queue := []Cell{start}
	visited[start] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v == end {
			return distance[end]
		}

		for _, adj := range adjacencies[v] {
			if checkValidCell(string(maps[adj.i][adj.j])) && !visited[adj] {
				visited[adj] = true
				distance[adj] = distance[v] + 1
				queue = append(queue, adj)
			}
		}
	}
	return 0
}

func checkValidCell(c string) bool {
	return c == "S" || c == "O" || c == "L" || c == "E"
}
