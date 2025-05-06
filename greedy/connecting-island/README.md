# [ 섬 연결하기 ](https://school.programmers.co.kr/learn/courses/30/lessons/42861)

경유하는 경로를 포함하여 모든 섬을 연결하는 다리를 지을 때 드는 최소비용을 구하는 문제

## 모범 풀이
가격 오름차순 정렬 후, 부모 노드를 재귀적으로 탐색하여 다리 건설 여부를 판별해야 함.

```go
type Sortable [][]int

func (s Sortable) Len() int           { return len(s) }
func (s Sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Sortable) Less(i, j int) bool { return (s[i][2] < s[j][2]) }

func solution(n int, costs [][]int) int {
	s := Sortable(costs)
	sort.Sort(s)
	costs = s[:]

	nodes := make(map[int][]int, 0)
	result := 0

	for i := 0; i < len(costs); i++ {
		if !traversal(nodes, make(map[int]bool, 0), costs[i][0], costs[i][1]) {
			nodes[costs[i][0]] = append(nodes[costs[i][0]], costs[i][1])
			nodes[costs[i][1]] = append(nodes[costs[i][1]], costs[i][0])
			result += costs[i][2]
		}
	}
	return result
}

func traversal(nodes map[int][]int, visit map[int]bool, start, end int) bool {
	if visit[start] {
		return false
	}

	visit[start] = true

	for _, v := range nodes[start] {
		if visit[v] {
			continue
		}

		if v == end {
			return true
		}

		if traversal(nodes, visit, v, end) {
			return true
		}
	}
	return false
}
```