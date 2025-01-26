# [ 이중 우선순위 큐 ](https://school.programmers.co.kr/learn/courses/30/lessons/42628?language=go)

## 모범 풀이
오퍼레이션에 따라 수를 추가하거나 현 시점의 최대,최솟값을 POP 한 뒤 최종적인 최대,최솟값을 반환하는 문제

물론 나처럼 풀 수 있겠지만 Insert마다 전체 배열을 Sorting하기 때문에 시간 비효율적임.
heap을 구현하여 처리하는 게 좋음

```go
import (
    "container/heap"
    "strings"
    "strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func solution(operations []string) []int {
    h := &IntHeap{}
    heap.Init(h)

    for _,op := range operations{
        items := strings.Split(op," ")
        number, _ := strconv.Atoi(items[1])
        length := h.Len()

        if items[0] == "I"{
            heap.Push(h, number);
        }else if length > 0 {
            if number == 1 {
                //최댓값 삭제
                heap.Remove(h, length-1)
            }else{
                //최솟값 삭제
                heap.Pop(h)
            }
        }
    }

    var ret = []int {0,0}

    checkIndex := 0
    //0 최대 // 1 최소값    
    for h.Len() > 0 {
        val := heap.Pop(h)

        if checkIndex == 0 {
            //First 최소값
            ret[1] = val.(int)
            checkIndex = 1
        }else if h.Len() == 0{
            //Last 최대값
            ret[0] = val.(int)
        }
    }

    return  ret
}
```