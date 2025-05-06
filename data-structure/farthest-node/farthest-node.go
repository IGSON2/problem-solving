package farthestnode

type Graph map[int][]int

var (
	tree     = make(Graph)
	maxDepth = 0
	deepest  = []node{}
	visited  = map[int]bool{}
)

type node struct {
	val   int
	depth int
}

func bfs() {
	queue := []node{{1, 0}}

	for len(queue) > 0 {

		n := queue[0]
		queue = queue[1:]

		visited[n.val] = true

		if n.depth > maxDepth {
			maxDepth = n.depth
			deepest = []node{n}
		} else if n.depth == maxDepth {
			deepest = append(deepest, n)
		}

		for _, child := range tree[n.val] {
			if !visited[child] {
				queue = append(queue, node{child, n.depth + 1})
			}
		}

	}
}

func buildTree(edges [][]int) {
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		tree[from] = append(tree[from], to)
		tree[to] = append(tree[to], from) // 무방향 그래프이므로 양방향 연결
	}
}

func solution(n int, edge [][]int) int {
	buildTree(edge)

	bfs()
	return len(deepest)
}
