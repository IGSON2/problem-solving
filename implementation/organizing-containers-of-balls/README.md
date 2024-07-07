# [Organaizing containers of balls](https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem?isFullScreen=true)

### n*n 이중배열이 주어졌을 때 행을 컨테이너, 열을  공의 타입이라고 가정한 뒤 오직 컨테너 간 공 교환만을 통해 한 컨테이너 마다 한 종류의 공이 담겨지도록 처리하는 문제

---
## 나의 풀이
분명 법칙이 있을거라 생각해봐도 도저히 답을 떠오르지 않아 포기함.

---

## 모범 풀이

```go
func main() {
    var q,n int
    fmt.Scanln(&q)
    for i := 0; i < q; i++ {
        fmt.Scanln(&n)
        row := make([]int,n)
        col := make([]int,n)
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                var ele int
                fmt.Scan(&ele)
                row[i] += ele
                col[j] += ele
            }
        }
        sort.Ints(row)
        sort.Ints(col)
        ok := true
        for i := 0; i < n; i++ {
            if (row[i] != col[i]) {
                ok = false
                break
            }
        }
        if (ok) {
            fmt.Println("Possible")
        } else {
            fmt.Println("Impossible")
        }
    }

}
```

문제에서 요구하는대로 컨테이너마다 한 종류의 공들을 모두 담아야 한다면 한 컨테이너 내의 모든 공의 합은 여러 컨테이너에 퍼져있는 특정 타입의 공의 합과 같아야 한다.