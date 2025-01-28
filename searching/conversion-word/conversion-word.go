package conversionword

func solution(begin string, target string, words []string) int {
	neighbor := make(map[string][]string)
	wordLength := len(begin)

	newWords := []string{begin}
	newWords = append(newWords, words...)

	for i := 0; i < len(newWords); i++ {
		wordA := newWords[i]

		for j := i + 1; j < len(newWords); j++ {
			wordB := newWords[j]
			compareCnt := 0
			for k := 0; k < wordLength; k++ {
				if wordA[k] == wordB[k] {
					compareCnt++
				}
			}
			if compareCnt == len(wordA)-1 {
				if neighbor[wordA] == nil {
					neighbor[wordA] = []string{wordB}
					continue
				}
				neighbor[wordA] = append(neighbor[wordA], wordB)
			}
		}
	}

	checked := make(map[string]bool)

	type depth struct {
		word  string
		depth int
	}
	queue := []depth{{begin, 0}}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if checked[v.word] {
			continue
		}

		checked[v.word] = true

		if v.word == target {
			return v.depth
		}

		for _, ele := range neighbor[v.word] {
			if !checked[ele] {
				queue = append(queue, depth{ele, v.depth + 1})
			}
		}
	}

	return 0
}
