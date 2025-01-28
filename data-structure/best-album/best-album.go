package bestalbum

import "container/heap"

type Item struct {
	no    int
	genre string
	play  int
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int           { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i].play > p[j].play }
func (p PriorityQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(*Item)) }
func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func solution(genres []string, plays []int) []int {
	p := new(PriorityQueue)
	heap.Init(p)

	for i := 0; i < len(genres); i++ {
		heap.Push(p, &Item{
			no:    i,
			genre: genres[i],
			play:  plays[i],
		})
	}

	result := make([]int, 0)
	checked := make(map[string][]int)
	getIdxByGenre := func(item *Item) int {
		idx, exist := checked[item.genre]
		if exist {
			return idx[0]
		}
		return -1
	}

	for p.Len() > 0 {
		item := heap.Pop(p).(*Item)
		if len(checked[item.genre]) > 1 {
			continue
		}
		idx := getIdxByGenre(item)
		if idx == -1 {
			result = append(result, item.no)
			checked[item.genre] = append(checked[item.genre], len(result)-1)
			continue
		}
		cp1 := make([]int, idx+1)
		cp2 := make([]int, len(result)-idx-1)
		copy(cp1, result[:idx+1])
		cp1 = append(cp1, item.no)
		copy(cp2, result[idx+1:])
		result = append(cp1, cp2...)
		checked[item.genre] = append(checked[item.genre], idx+1)
	}

	return result
}
