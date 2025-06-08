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
