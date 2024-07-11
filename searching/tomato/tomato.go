package tomato

import (
	"bufio"
	"fmt"
	"os"
	"problem-solving/utils"
)

type spot struct {
	i, j int
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var moves = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func tomato() {
	f, err := os.Open("testcase")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var W, H int
	fmt.Fscanf(f, "%d %d\n", &W, &H)
	queue := make([]spot, 0)

	box := make([][]int16, H)
	for i := 0; i < H; i++ {
		box[i] = append(box[i], make([]int16, W)...)
		for j := 0; j < W; j++ {
			var t int8
			fmt.Fscan(f, &t)
			box[i][j] = int16(t)
			if t == 1 {
				queue = append(queue, spot{i, j})
			}
		}
	}
	// bfs
	visited := make(map[spot]bool)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, m := range moves {
			ni, nj := v.i+m[0], v.j+m[1]
			if !isSafe(ni, nj, W, H) || visited[spot{ni, nj}] || box[ni][nj] == -1 {
				continue
			}
			cur := box[v.i][v.j]
			box[ni][nj] = cur + 1
			queue = append(queue, spot{ni, nj})
		}
		visited[v] = true
	}

	var result int16
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if box[i][j] == 0 {
				fmt.Println(-1)
				return
			}

			if result < box[i][j] {
				result = box[i][j]
			}
		}
	}
	utils.PrintMemUsage()
	fmt.Println(result - 1)
}

func isSafe(i, j, W, H int) bool {
	return i >= 0 && j >= 0 && i < H && j < W
}
