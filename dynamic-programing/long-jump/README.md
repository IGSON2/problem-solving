# [멀리뛰기](https://school.programmers.co.kr/learn/courses/30/lessons/12914)

### 예를 들어, n개의 계단이 있다고 가정하면, 이때 n번째 계단에 도달하는 방법
* n-1번째 계단에서 1칸 올라오는 경우
* n-2번째 계단에서 2칸 올라오는 경우

이 두 가지 경우로 나눌 수 있다. 따라서, n번째 계단까지 올라가는 방법의 수는 n-1번째 계단까지 올라가는 방법의 수와 n-2번째 계단까지 올라가는 방법의 수를 더한 것과 같다.

```go
func solution(n int) int64 {

   var dp [2001]int
    dp[0]=1
    dp[1]=1
    for i:=2; i<=n; i++ {
        dp[i]= (dp[i-2] + dp[i-1])%1234567
    }


    return int64(dp[n])
}
```

### n개중 무작위로 r개를 뽑는 경우의 수   
`C(n,r)= n! / r!×(n−r)!`
​
