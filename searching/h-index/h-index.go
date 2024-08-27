package hindex

import "sort"

func solution(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	h := citations[0]
	for h > 0 {
		if sort.Search(len(citations), func(i int) bool { return citations[i] < h }) >= h {
			return h
		}
		h--
	}
	return h
}
