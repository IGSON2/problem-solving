# [ 택배상자 주문 순서에 맞춰 처리하기 ](https://school.programmers.co.kr/learn/courses/30/lessons/131704)

## 모범 풀이
큐로 들어오는 데이터를 처리하고, 처리하지 못한 데이터는 스택에 적재하여 한번 더 처리하는 방식으로 풀이.

```go
func betterSolution(order []int) int {
	answer := 0
	var s []int
    for i := 1; i <= len(order); i++ {
		if i == order[answer] {
			answer++
		} else {
			s = append(s, i)
		}
		for 0 < len(s) && s[len(s)-1] == order[answer] {
			answer++
			s = s[:len(s)-1]
		}
	}
	return answer
}
```