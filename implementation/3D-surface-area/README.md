# [3D surface area](https://www.hackerrank.com/challenges/3d-surface-area/problem)

### 3차원 구조물의 표면적을 구하는 문제

---

## 1차 풀이

```go
func main() {
    var H,W int
    fmt.Scanf("%d %d",&H,&W)
    rowHighest := make([]int,H)
    colHighest := make([]int,W)
    
    for i := 0; i < H; i++{
        for j := 0; j < W; j++{
            var tmp int
            fmt.Scan(&tmp)
            
            if tmp > rowHighest[i]{
                rowHighest[i] = tmp
            }
            
            if tmp > colHighest[j]{
                colHighest[j] = tmp
            }
        }
    }
    
    sum := func(nums []int)int{
        var sum int
        for _,n := range nums{
            sum += n
        }
        return sum
    }
    result := 2 * sum(rowHighest) + 2 * sum(colHighest) + 2 * W*H
    fmt.Println(result)
}
```
한 측면에서 바라봤을 때 가장 높은 높이들을 구하여 더하면 될 줄 알았지만, [3,2,3]과 같이 높이가 같은 구역 사이에 높이가 낮은 구역이 위치할 경우 새로운 표면이 생기는 현상을 고려하지 못했다.

## 2차 풀이
```go
func Abs(a int)int{
    if a < 0{
        return -a
    }
    return a
}

func main() {
    var H,W int
    fmt.Scanf("%d %d",&H,&W)
    toy := make([][]int,H)
    
    for i := 0; i < H; i++{
        for j := 0; j < W; j++{
            var tmp int
            fmt.Scan(&tmp)
            toy[i] = append(toy[i], tmp)
        }
    }
    
    result := 2*W*H // 밑면 + 윗면
    for i := 0; i < H; i++{
        for j := 0; j < W; j++{
            h := toy[i][j]
            if i == 0 {
                result += h
            }
            if j == 0 {
                result += h
            }
            // i == 0 && j == 0 인 경우(모서리)는 h가 두 번 더해짐
            
            if i == H-1 { 
                result += h
            }else{
                result += Abs(toy[i][j]-toy[i+1][j])
            }
            
            if j == W-1{
                result += h
            }else{
                result += Abs(toy[i][j]-toy[i][j+1])
            }
        }
    }
    
    fmt.Println(result)
}
```