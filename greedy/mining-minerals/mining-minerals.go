package miningminerals

import (
	"math"
	"sort"
)

func MiningMinerals(picks []int, minerals []string) int {
	result := 0

	power := 25
	nPicks := make([]int, 0)
	for _, p := range picks {
		for i := 0; i < p; i++ {
			nPicks = append(nPicks, power)
		}
		power /= 5
	}

	N := len(minerals)/5 + 1
	if len(minerals)%5 == 0 {
		N--
	}

	if len(minerals) > len(nPicks)*5 {
		minerals = minerals[:len(nPicks)*5]
	}

	mine := make([][]int, N)
	for i := 0; i < N; i++ {
		mine[i] = make([]int, 6)
	}

	sum := 0
	for i, m := range minerals {
		fatigue := 0

		switch m {
		case "diamond":
			fatigue = 25
		case "iron":
			fatigue = 5
		case "stone":
			fatigue = 1
		}
		mine[i/5][i%5] = fatigue
		sum += fatigue

		if (i+1)%5 == 0 || i == len(minerals)-1 {
			mine[i/5][5] = sum
			sum = 0
		}
	}
	sort.Slice(mine, func(i, j int) bool {
		return mine[i][5] > mine[j][5]
	})

	i := 0
	for len(nPicks) != 0 && i < len(mine) {
		p := nPicks[0]
		nPicks = nPicks[1:]

		for j := 0; j < 5; j++ {
			result += int(math.Ceil(float64(mine[i][j]) / float64(p)))
		}
		i++
	}
	return result
}
