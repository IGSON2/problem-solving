package journeytothemoon

func dfs(graph map[int32][]int32, visited map[int32]bool, node int32) int32 {
	stack := []int32{node}
	var size int32

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[node] {
			visited[node] = true
			size++

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	return size
}

func journeyToMoon(n int32, astronaut [][]int32) int32 {
	graph := make(map[int32][]int32)
	for _, a := range astronaut {
		graph[a[0]] = append(graph[a[0]], a[1])
		graph[a[1]] = append(graph[a[1]], a[0])
	}

	visited := make(map[int32]bool)
	pairs := make([]int32, 0)
	for i := 0; i < int(n); i++ {
		pairs = append(pairs, dfs(graph, visited, int32(i)))
	}

	var result int32
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			result += pairs[i] * pairs[j]
		}
	}

	return result
}
