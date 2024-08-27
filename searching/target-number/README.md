# [Target Number](https://school.programmers.co.kr/learn/courses/30/lessons/43165)

# SEXY CODE
```go
func solution(numbers []int, target int) int {
    if len(numbers) == 0 && target == 0 {
        return 1
    } else if len(numbers) == 0 {
        return 0
    }
    return solution(numbers[1:],target-numbers[0])+ solution(numbers[1:],target+numbers[0])    
}
```