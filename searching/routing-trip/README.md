# [여행 경로](https://school.programmers.co.kr/learn/courses/30/lessons/43164?language=go)

나의 풀이로는 반례 `tickets := [][]string{{"ICN", "AAA"}, {"ICN", "BBB"}, {"BBB", "ICN"}}`를 해결하지 못하였음
해결 불가 원인
- 경로가 막히면 백트래킹하지 않고 종료됨
- 티켓이 모두 사용되었는지도 확인하지 않았음.

### 나의 풀이
```go
package routingtrip

import (
	"sort"
)

const ICN = "ICN"

func solution(tickets [][]string) []string {
	routes := make(map[string][]string)
	for _, t := range tickets {
		routes[t[0]] = append(routes[t[0]], t[1])
	}

	for _, v := range routes {
		sort.Strings(v)
	}

	return dfs(routes, ICN)
}

func dfs(routes map[string][]string, from string) []string {
	result := []string{from}

	if len(routes[from]) > 0 {
		r := routes[from][0]
		routes[from] = routes[from][1:]
		result = append(result, dfs(routes, r)...)
	}
	return result
}

```