# [ 주어진 숫자 배열을 문자화 하여 조합하였을 때 나올 수 있는 가장 큰 수 ](https://school.programmers.co.kr/learn/courses/30/lessons/42746)

## 모범 풀이
삽질만 함.
자릿수가 작은 수가 무조건 앞으로 와야 한다고 착각함. [3,301] 같은 경우만 생각함.
그러나 [3,34] 같은 경우도 있기 때문에 모든 자릿수를 비교했어야 함.
방법을 아예 찾지 못했음

```go
func solution(numbers []int) string {
    sort.Slice(numbers, func(i, j int) bool {
        s1 := strconv.Itoa(numbers[i])
        s2 := strconv.Itoa(numbers[j])
        return s1+s2 > s2+s1 // 3이 301보다 앞으로 정렬됨
    })

    if numbers[0] == 0 {
        return "0"
    }
    answer := ""
    for i := range numbers {
        answer += strconv.Itoa(numbers[i])
    }
    return answer
}
```

문자열을 조합하여 정렬하는 방법이 있을거라곤 생각도 못함.
"3301" > "3013"