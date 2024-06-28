package pairs

// 이중 반복문을 이용한 탐색 -> O(n^2)
func pairs(k int32, arr []int32) int32 {
	var cnt int32
	for i, a := range arr {
		var tempCnt int32
		for j := i + 1; j < len(arr); j++ {
			if sub := a - arr[j]; sub == k || sub == -k {
				tempCnt++
				if a < k {
					break
				}
			}
			if tempCnt == 2 {
				break
			}
		}
		cnt += tempCnt
	}
	return cnt
}

// Hash맵을 이용한 탐색 -> O(n)
func pairs2(k int32, arr []int32) int32 {
	var cnt int32
	checked := make(map[int32]struct{})
	for _, a := range arr {
		checked[a] = struct{}{}
	}

	for _, a := range arr {
		if _, exist := checked[a+k]; exist {
			cnt++
		}

		if _, exist := checked[a-k]; exist {
			cnt++
		}
		delete(checked, a) // 중복 제거
	}
	return cnt
}
