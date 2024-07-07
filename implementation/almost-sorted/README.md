# [Almost sorted](https://www.hackerrank.com/challenges/almost-sorted/problem)

### 정렬된 배열을 만들기 위해 swap 또는 reverse를 한번만 사용하여 정렬할 수 있는지 판단하는 문제

---


## 내 풀이
```go
func almostSorted(arr []int32) {
    indices := make([]int,0)
    for i := 0; i < len(arr)-1; i++{
        if arr[i] > arr[i+1]{
            indices = append(indices, i)
        }
    }
    
    if len(indices) == 0 {
        fmt.Println("yes")
        return
    }
    
    if src := arr[indices[0]]; len(indices) == 1{
        for i := indices[0]+2; i < len(arr); i++{
            if arr[i] > arr[indices[0]+1]
        }
    }
    
    subArr := make([]int32,0)
    copy(subArr, arr[indices[0]:indices[1]+1])
    
    for i := 1; i < len(subArr); i++ {
        if subArr[i-1] < subArr[i]{
            fmt.Println("no")
        }
    }
    
    if arr[indices[0]-1] < arr[indices[1]] && arr[indices[0]] < arr[indices[1]+1]{
        fmt.Printf("yes\nreverse %d %d",indices[0]+1,indices[1]+1)
    }
    fmt.Println("error")
}
```
감이 안잡혀 산으로 가는 기분이 들어 포기하였다. 핵심은 정렬에 부합하지 않는 요소의 갯수에 있었다.


## 모범 답안
```go
func check(orig, arr, diff []int) bool {
    for i := 1; i < len(diff); i++ {
        if orig[diff[i-1]] < orig[diff[i]] { return false }
        if diff[i] - diff[i-1] != 1 { 
            for j := diff[i-1]+1; j < diff[i]; j++ {
                if orig[j] != arr[j] { return false }
            } 
        }
    }
    return true
}

func f(arr []int) {
    orig := make([]int, len(arr))
    copy(orig, arr)
    sort.Ints(arr)
    diff := []int{}
    for i := 0; i < len(arr); i++ {
        if arr[i] != orig[i] {
            diff = append(diff, i)
        }
    }
    if len(diff) == 0 {
        fmt.Println("yes")
    } else if len(diff) == 2 {
        fmt.Println("yes")
        fmt.Println("swap", diff[0]+1, diff[1]+1)
    } else if len(diff) > 2 {
        if check(orig, arr, diff) {
            fmt.Println("yes")
            fmt.Println("reverse", diff[0]+1, diff[len(diff)-1]+1)    
        } else {
            fmt.Println("no")
        }        
    }else {
        fmt.Println("no")
    }
}
```