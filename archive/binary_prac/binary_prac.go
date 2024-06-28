package binaryprac

func binarySearch(mid int, ranked []int32, highScore, pScore int32) int32 {
	var result int32 = ranked[mid]
	if result < highScore {
		highScore = result
	}

	if pScore >= result {
		binarySearch(mid/2, ranked, highScore, pScore)
	} else {
		binarySearch(mid+(len(ranked)-mid)/2, ranked, highScore, pScore)
	}

	return highScore
}
