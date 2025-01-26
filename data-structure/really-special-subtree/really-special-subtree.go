package reallyspecialsubtree

import (
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func reallySpecialSubtree() {
	f, err := os.Open("testcase")
	if err != nil {
		panic(err)
	}
	var gNodes, gEdges int
	fmt.Fscanf(f, "%d %d", &gNodes, &gEdges)

	G := make(map[int]map[int]int)
	for i := 0; i < gEdges; i++ {
		var edgeFrom, edgeTo, edgeWeight int
		fmt.Fscanf(f, "%d %d %d", &edgeFrom, &edgeTo, &edgeWeight)

		if len(G[edgeFrom]) == 0 {
			G[edgeFrom] = make(map[int]int)
		}
		if len(G[edgeTo]) == 0 {
			G[edgeTo] = make(map[int]int)
		}

		val, ok := G[edgeFrom][edgeTo]
		if ok {
			if edgeWeight < val {
				G[edgeFrom][edgeTo] = edgeWeight
			}
		} else {
			G[edgeFrom][edgeTo] = edgeWeight
		}

		val, ok = G[edgeTo][edgeFrom]
		if ok {
			if edgeWeight < val {
				G[edgeTo][edgeFrom] = edgeWeight
			}
		} else {
			G[edgeTo][edgeFrom] = edgeWeight
		}
	}
	weight := prim(G, gNodes, 1)
	fmt.Println(weight)
}

func prim(G map[int]map[int]int, nodes int, start int) int {
	connected := make([]bool, nodes+1)
	var num, res int
	frontier := make(PriorityQueue, 1)
	frontier[0] = &Item{value: start, priority: 0, index: 0}
	heap.Init(&frontier)

	for frontier.Len() > 0 && num <= nodes {
		element := heap.Pop(&frontier).(*Item)
		vertex := element.value

		if !connected[vertex] {
			res += element.priority
			num++
			connected[vertex] = true

			for v, w := range G[vertex] {
				if !connected[v] {
					heap.Push(&frontier, &Item{value: v, priority: w})
				}
			}
		}
	}
	return res
}
