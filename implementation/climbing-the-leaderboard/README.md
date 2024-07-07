# [Climbing the leaderboard](https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem)

### Rank 배열과 Player 배열이 주어졌을 때 각각의 Player가 어떤 순위에 위치하는지를 반환하는 문제이다.

---

## 나의 풀이
```go
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    discinctRanks := make([]int32, 1)
    discinctRanks[0] = -1
    
    // 중복 제거, 1등의 인덱스가 [1]임에 유의
    for _,r := range ranked{
        if discinctRanks[len(discinctRanks)-1] == r{
            continue
        }
        discinctRanks = append(discinctRanks, r)
    }
    result := make([]int32,0)
    
    var cache int = 1
    Outer:
    for i := len(player)-1; i >= 0; i-- {
        for j := cache; j < len(discinctRanks); j++{
            if player[i] >= discinctRanks[j]{
                result = append([]int32{int32(j)},result...)
                cache = j
                continue Outer
            }
        }
        result = append([]int32{int32(len(discinctRanks))},result...)
    }
    
    return result
}
```

12개의 케이스 중 Case 8 에서 Time limit exceeded에 걸렸다.

---

## 모범 풀이
```go
func main() {
	n := readInt()
	distinctScores := make([]int, 0, n)
	// Don't really need the count, but why not
	scoreCounts := make(map[int]int)

	for i := 0; i < n; i++ {
		x := readInt()
		if scoreCounts[x] == 0 {
			distinctScores = append(distinctScores, x)
		}
		scoreCounts[x]++
	}

	m := readInt()
	for i := 0; i < m; i++ {
		x := readInt()
		// Find the rank of this score
		idx := sort.Search(len(distinctScores), func(i int) bool {
			return distinctScores[i] <= x
		})
		fmt.Println(idx + 1)
	}

}

func readInt() (x int) {
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		panic(err)
	}
	return
}
```

`sort.Search()`를 사용하여 풀이하였다. `sort.Search()`가 탐색시간이 더 빠른걸까?