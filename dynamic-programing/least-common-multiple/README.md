# [최소 공배수](https://school.programmers.co.kr/learn/courses/30/lessons/12953)

### 숫자 배열 arr의 모든 원소에 대한 최소 공배수를 구하는 문제

## 모범 풀이
```go
func gcd(a, b int)int{
    if a%b == 0{
        return b
    }else{
        return gcd(b, a%b)
    }
}

func lcm(nums []int)int{
    ans:=1
    for _, num := range nums{
        ans *= num/gcd(ans, num)
    }
    return ans
}

func solution(arr []int) int {
    return lcm(arr)
}
```

### 최대 공약수와 최소 공배수의 관계
최대 공약수 G와 최소 공배수 L을 곱한 `G x L`은 두 수 `A x B`와 같다.   
두 수를 G로 나눈 `a, b`에 대해 `a x b x G`는 `L`과 같다.