package bestalbum

import "sort"

type Item struct {
	no   int
	play int
}

func solution(genres []string, plays []int) []int {
	playCnt := make(map[string]int)
	itemsByGenres := make(map[string][]*Item)
	for i := 0; i < len(genres); i++ {
		itemsByGenres[genres[i]] = append(itemsByGenres[genres[i]], &Item{i, plays[i]})
		playCnt[genres[i]] += plays[i]
	}

	for k, v := range itemsByGenres {
		sorted := make([]*Item, len(v))
		copy(sorted, v)
		sort.Slice(sorted, func(i, j int) bool {
			if sorted[i].play == sorted[j].play {
				return sorted[i].no < sorted[j].no
			}
			return sorted[i].play > sorted[j].play
		})
		itemsByGenres[k] = sorted
	}

	sortedGenres := make([]string, 0)
	for len(playCnt) > 0 {
		topGenre := ""
		max := 0
		for k, v := range playCnt {
			if v > max {
				max = v
				topGenre = k
			}
		}
		sortedGenres = append(sortedGenres, topGenre)
		delete(playCnt, topGenre)
	}

	result := make([]int, 0)
	for _, genre := range sortedGenres {
		v := itemsByGenres[genre]
		for i := 0; i < len(v) && i < 2; i++ {
			result = append(result, v[i].no)
		}
	}

	return result
}
