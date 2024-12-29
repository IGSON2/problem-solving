package uninhabitedislandtrip

import (
	"sort"
	"strconv"
	"strings"
)

type Cell struct {
	i, j int
}

func (c *Cell) isSafe(M, N int) bool {
	return c.i >= 0 && c.j >= 0 && c.i < M && c.j < N
}

var direction = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solution(maps []string) []int {
	result := make([]int, 0)
	M, N := len(maps), len(maps[0])
	map2D := make([][]string, M)

	for i, m := range maps {
		line := strings.Split(m, "")
		map2D[i] = line
	}

	adj := make(map[Cell][]Cell)
	for i, m := range map2D {
		for j, m2 := range m {
			if m2 == "X" {
				continue
			}
			c := Cell{i: i, j: j}

			for _, d := range direction {
				nc := Cell{c.i + d[0], c.j + d[1]}
				if nc.isSafe(M, N) {
					adj[c] = append(adj[c], nc)
				}
			}
		}
	}

	visited := make(map[Cell]bool)
	for i, m := range map2D {
		for j, m2 := range m {
			if m2 != "X" && !visited[Cell{i, j}] {
				r := dfs(Cell{i, j}, adj, visited, map2D, 0)
				if r > 0 {
					result = append(result, r)
				}
			}
		}
	}

	if len(result) == 0 {
		return []int{-1}
	}

	sort.Ints(result)
	return result
}

func dfs(start Cell, adj map[Cell][]Cell, visited map[Cell]bool, map2D [][]string, cnt int) int {
	visited[start] = true
	f, _ := strconv.Atoi(map2D[start.i][start.j])
	cnt += f

	for _, a := range adj[start] {
		if !visited[a] && map2D[a.i][a.j] != "X" {
			visited[a] = true
			cnt = dfs(a, adj, visited, map2D, cnt)
		}
	}
	return cnt
}
