package farthestnode

import (
	"container/heap"
	"math"
)

type Heap []Cost

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Cost))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

type Edge struct {
	from   int
	to     int
	weight int
}

type Cost struct {
	to   int
	cost int
}

func HeapSolution(n int, edge [][]int) int {
	adj := [20001][]*Edge{}
	costs := [20001]int{}

	for i := 2; i <= n; i++ {
		costs[i] = math.MaxInt64
	}

	for _, e := range edge {
		from := e[0]
		to := e[1]
		adj[from] = append(adj[from], &Edge{from: from, to: to, weight: 1})
		adj[to] = append(adj[to], &Edge{from: to, to: from, weight: 1})
	}

	h := &Heap{Cost{to: 1, cost: 0}}
	heap.Init(h)

	for (*h).Len() != 0 {
		e := h.Pop().(Cost)
		if costs[e.to] < e.cost {
			continue
		}

		for _, v := range adj[e.to] {
			if costs[v.to] > costs[e.to]+v.weight {
				costs[v.to] = costs[e.to] + v.weight
				h.Push(Cost{to: v.to, cost: costs[v.to]})
			}
		}
	}

	max := 0
	maxCount := 0
	for _, v := range costs {
		if v > max {
			max = v
			maxCount = 1
		} else if v == max {
			maxCount++
		}
	}

	return maxCount
}
