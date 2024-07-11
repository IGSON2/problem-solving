# [Grid Search](https://www.hackerrank.com/challenges/the-grid-search/problem)

### 큰 이중 배열에 작은 이중 배열이 포함되어 있는지 확인하는 문제이다.

---

## 나의 풀이

```go
func gridSearch(G []string, P []string) string {
    var answer string = "NO"
    var pLen int = len(P[0])
    
    var wPattern string
    for _, p := range P{
        wPattern += p
    }
    
    for i := 0; i <= len(G[0])-pLen; i++{
        var c string
        for j := 0; j < len(G); j++{
            c += G[j][i:i+pLen]
        }
        if strings.Contains(c,wPattern){
            answer = "YES"
            break
        }    
    }
    
    return answer
}
```
처음에는 그리드와 패턴을 비교할 때 인덱스 하나하나 비교하려고 했는데, 열을 기준으로 문자열을 묶어 비교하는 방법이 더 간단했다.
그러나 이 방법은 시간복잡도가 O(n^2)이기 때문에 더 효율적인 방법이 있을 것이다.

## 슬라이딩 윈도우를 이용한 풀이
```go
func gridSearch(G []string, P []string) string {
    rowsG, colsG := len(G), len(G[0])
    rowsP, colsP := len(P), len(P[0])
    
    // 그리드 G를 순회하기 시작. 단, 남은 탐색 row가 rowsP 보다 적게 남았을 땐 더이상 순회할 필요가 없다.
    for i := 0; i <= rowsG - rowsP; i++ {
        for j := 0; j <= colsG - colsP; j++ {   // 매 그리드 G의 row를 linear search한다. 단, 남은 탐색 column이 colsP 보다 적게 남았을 땐 더이상 순회할 필요가 없다.
            match := true
            for k := 0; k < rowsP; k++ {
                if G[i+k][j:j+colsP] != P[k] {
                    match = false
                    break
                }
            }
            if match {
                return "YES"
            }
        }
    }
    return "NO"
}
```

## 예시
```
    G := []string{
        "2229505",
        "5633845",
        "6473530",
        "7053106",
        "0834282",
        "4607924",
    }

    P := []string{
        "9505",
        "3845",
        "3530",
    }
```

```go
// j == 0
[2229]505 // => break
5633845
6473530
7053106
0834282
4607924

// j == 1
2[2295]05 // => break
5633845
6473530
7053106
0834282
4607924

...

// j == 3, k == 0
222[9505] // => match
5633845
6473530
7053106
0834282
4607924

...

// j == 3, k == 2
222[9505] // => match
563[3845] // => match
647[3530] // => match
7053106
0834282
4607924
```