# [ 의상 조합하기 ](https://school.programmers.co.kr/learn/courses/30/lessons/42578)

## 모범 풀이

```go
func solution(clothes [][]string) int {
    m := make(map[string]int)
    for _, c := range clothes {
        m[c[1]]++
    }
    answer := 1
    for _, v := range m {
        answer *= (v + 1)
    }
    return answer - 1
}
```

상의가 2개, 하의가 2개가 있다고 가정하면, 상의를 입는 경우의 수는 3가지(입거나, 입지 않거나)이고, 하의를 입는 경우의 수도 3가지이다.    
이후 상의를 입는 경우의 수와 하의를 입는 경우의 수를 곱하면 된다. 그리고 모두 입지 않는 경우는 없으므로 1을 빼주면 된다.