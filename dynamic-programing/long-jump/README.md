# [멀리뛰기](https://school.programmers.co.kr/learn/courses/30/lessons/12914)

피보나치를 이용해 푸는 게 효율적임
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
