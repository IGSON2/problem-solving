package ricochetbot

import (
	"fmt"
	"strings"
)

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type cell struct {
	i, j int
}

var nilCell = cell{i: -1, j: -1}

func solution(board []string) int {
	var start, goal cell

	adjacencies := make(map[cell][]cell)

	N, M := len(board), len(board[0])

	for i := 0; i < N; i++ {
		line := strings.Split(board[i], "")
		for j := 0; j < M; j++ {
			c := cell{i, j}

			if line[j] == "G" {
				goal = c
			} else if line[j] == "R" {
				start = c
			} else if line[j] == "D" {
				continue
			}

			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if isSafe(ni, nj, N, M) {
					if string(board[ni][nj]) != "D" {
						nc := cell{ni, nj}
						adjacencies[c] = append(adjacencies[c], nc)
						continue
					}
				}
				adjacencies[c] = append(adjacencies[c], nilCell)
			}
		}
	}

	wallsByCell := make(map[cell][4]cell)
	for k, v := range adjacencies {
		walls := [4]cell{nilCell, nilCell, nilCell, nilCell}
		for i := 0; i < 4; i++ {
			latest := nilCell
			adj := v
			for {
				cur := adj[i]
				if cur == nilCell {
					break
				}
				latest = cur
				adj = adjacencies[cur]
				if adj == nil {
					break
				}
			}
			walls[i] = latest
		}
		wallsByCell[k] = walls
	}

	visited := make(map[cell]int)

	moves := bfs(start, goal, &wallsByCell, visited)
	return moves
}

func isSafe(ni, nj, N, M int) bool {
	return ni >= 0 && nj >= 0 && ni < N && nj < M
}

func bfs(start, goal cell, walls *map[cell][4]cell, visited map[cell]int) int {
	queue := []cell{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		fmt.Println(v, visited[v])
		if v.i == goal.i && v.j == goal.j {
			return visited[v]
		}

		for _, w := range (*walls)[v] {
			if w != nilCell && visited[w] == 0 {
				visited[w] = visited[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return -1
}
