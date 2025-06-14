package routingtrip

import (
	"sort"
)

func solution(tickets [][]string) []string {
	n := len(tickets)
	visited := make([]bool, n)
	var answer []string
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})

	var dfs func(path []string, depth int)
	dfs = func(path []string, depth int) {
		if len(answer) > 0 {
			return // 이미 정답 찾았으면 종료
		}
		if depth == n {
			answer = append([]string{}, path...)
			return
		}

		for i, ticket := range tickets {
			if !visited[i] && path[len(path)-1] == ticket[0] {
				visited[i] = true
				dfs(append(path, ticket[1]), depth+1)
				visited[i] = false
			}
			// 라우팅 잘못 들었을 때 여기서 백트래킹됨
		}
	}

	dfs([]string{"ICN"}, 0)
	return answer
}
