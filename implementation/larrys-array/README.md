# [Larry's array](https://www.hackerrank.com/challenges/larrys-array/problem)

### 길이가 n인 배열 내에 1부터 n까지의 수가 중복 없이 들어있을 때, 정렬되지 않은 구간을 ABC -> BCA와 같이 인덱스를 스왑하여 정렬해 나갈수 있는지 판별하는 문제

---

## 내 풀이
```go
func swap(arr []int32, n int){
    arr[n],arr[n+2] = arr[n+2],arr[n]
    arr[n],arr[n+1] = arr[n+1],arr[n]
}

func larrysArray(A []int32) string {
    orig := make([]int32,len(A))
    copy(orig,A)
    Outer:
    for i := 0; i < len(orig)-2; i++ {
        if orig[i] != int32(i+1) {
            for j := 0; j < 3; j++ {
                swap(orig,i)
                fmt.Printf("idx : [%d] - %v\n",i,orig)
                if orig[i] == int32(i+1) {
                    continue Outer
                }
            }
        }
    }
    
    for i := 0; i < len(orig); i++{
        if int32(i+1) != orig[i]{
            return "NO"
        }
    }
    
    return "YES"
}
```
단순히 직선상으로 스왑해 나가는 케이스만 생각했다. 예를들어 `[1,6,5,2,4,3]`와 같은 케이스는 양 방향에서 복합적으로 생각해야하는 문제였는데 간과했던 것.
그런데 이 케이스는 문제 설명에도 나와있었기에 결국엔 내가 문제를 제대로 인지하지 못해 생긴 문제였다.

## 모범 풀이

```go
func permutation_sign(arr []int) int {
	total := 0
	for i := 1; i < len(arr) + 1; i++ {
		for j := i + 1; j < len(arr) + 1; j++ {
			if arr[i - 1] > arr[j - 1] {
				total += 1
			}
		}
	}
	if total % 2 == 1 {
		return -1
	}
	return 1
}

func main(){
	T, reader := 0, bufio.NewReader(os.Stdin)
    fmt.Scanf("%d", &T)
    for ; T > 0; T--{
		_, _ = reader.ReadString('\n')
		text, _ := reader.ReadString('\n')
		arr := []int{}
		for _, v := range (strings.Fields(text)) {
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
		if permutation_sign(arr) == 1 {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
```

코드를 보고 이해가 가지 않았다.  
왜 이중 루프를 돌며 순열에 어긋나있는 부분을 카운트하고 그 수가 2로 나누어지면 정렬 가능하다고 확정해버리는걸까?  

### 홀순열과 짝순열
정렬되지 않은 순열을 정렬하기 위해 필요한 교환 횟수에 따라 순열을 두 가지 유형으로 나눌 수 있다.

1. 짝순열 (Even Permutation): 순열을 정렬하기 위해 필요한 교환 횟수가 짝수인 순열이다. 예를 들어, [1, 2, 3]은 이미 정렬된 상태이므로 교환 횟수는 0이다. 0은 짝수이므로 이는 짝순열이다.  
2. 홀순열 (Odd Permutation): 순열을 정렬하기 위해 필요한 교환 횟수가 홀수인 순열이다. 예를 들어, [2, 1, 3]은 정렬하기 위해 [1, 2, 3]으로 변경해야 하며, 1번의 교환이 필요하다. 1은 홀수이므로 이는 홀순열이다.

여기서 문제에서 제시한 회전연산 `ABC -> BCA -> CAB -> ABC` 중 `ABC -> BCA`를 자세히 들여다보면 아래처럼 2회 교환이 일어난다는 것을 알 수 있다.

```
1. A와 C를 교환 -> CBA
2. C와 B를 교환 -> BCA
```
3개의 요소를 회전시키는 연산은 항상 두 번의 교환으로 이루어지며, 두 번의 교환은 짝수 번이다. 이 점이 중요한 이유는 다음과 같다

1. 짝수 번의 교환은 순열 부호에 영향을 미치지 않음
   * 짝수 번의 교환은 순열의 부호를 변경하지 않는다. 즉, 짝순열은 여전히 짝순열로 유지되고, 홀순열은 여전히 홀순열로 유지된다.  

2. 홀수 번의 교환은 순열 부호를 변경함
   * 반대로, 홀수 번의 교환은 순열의 부호를 변경한다. 홀순열은 짝순열로, 짝순열은 홀순열로 바뀐다.
  
**즉 짝순열은 짝수번의 교환으로, 홀순열은 홀수번의 교환으로 정렬 가능하단 의미이며 주어진 배열이 짝순열인지만 판별하면 된다는 것.**

### 순열 유형 파악하기
**역전 쌍**은 앞에 있는 원소가 뒤에 있는 원소보다 큰 경우를 의미한다.  
이 역전쌍의 총 갯수가 홀수이면 홀순열, 짝수이면 짝순열이 된다.

```go
// 모범 답안에서 순열 유형을 파악하는 로직
func permutation_sign(arr []int) int {
	total := 0
	for i := 1; i < len(arr) + 1; i++ {
		for j := i + 1; j < len(arr) + 1; j++ {
			if arr[i - 1] > arr[j - 1] {
				total += 1
			}
		}
	}
	if total % 2 == 1 {
		return -1
	}
	return 1
}
```