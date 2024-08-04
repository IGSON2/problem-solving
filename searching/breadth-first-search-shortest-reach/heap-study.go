package breadthfirstsearchshortestreach

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value, priority, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func PrintHeap() {
	pq := make(PriorityQueue, 0)
	for i := 1; i < 4; i++ {
		pq = append(pq, &Item{
			value:    i * 10,
			index:    i,
			priority: 10 - i*2,
		})
	}
	for _, v := range pq {
		fmt.Println(v)
	}
	heap.Init(&pq)
	fmt.Println()
	heap.Push(&pq, &Item{value: 100, index: 5, priority: 0})
	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}
}
