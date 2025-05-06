# [ 가장 먼 노드 ](https://school.programmers.co.kr/learn/courses/30/lessons/49189?language=go)

간선이 있는 노드 군집에서 1번 노드로부터 제일 멀리 떨어져있는 노드들의 갯수를 구하는 문제

## 나의 풀이
하위 인스턴스를 참조하는 타입을 선언하지 않고 인덱스를 값으로 가지는 타입만 선언하여 BFS로 해결하였다.  
BFS를 모두 마칠 때 까지 최대 Depth를 저장하는 배열을 갱신하고 추가하도록 구현하였다.  
child 노드를 큐에 추가하기 전 visited 체크를 해 주어야 세 개 이상의 간선을 가진 노드 탐색의 오류를 방지할 수 있었다.

## 모범 풀이

이 문제의 의도는 힙을 통해 구현한 우선순위 큐로 “간선의 코스트가 1인 다익스트라 알고리즘을 구현할 수 있는가” 였을 것이다.  
heap은 그래프 자료구조이며, Go에선 heap 인터페이스를 구현해야 내장 패키지를 통해 편리하게 사용할 수 있다.  
여기서 관건은 어떤 값에 가중치를 둘 것이며 정렬 기준은 어떻게 정할 것인지 스스로 판단해야 하는 것인데, 이 부분이 미숙하여 편한 방법으로 풀게 됐다.  
아래처럼 풀었다면 스스로도 만족했을 것 같다.


```go
import "container/heap"
import "math"

type Heap []Cost

func (h Heap) Len() int {return len(h)}
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h Heap) Less(i, j int) bool {return h[i].cost< h[j].cost}
func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(Cost))
}

func (h *Heap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[0:len(*h)-1]
    return x
}

type Edge struct {
    from int
    to int
    weight int
}

type Cost struct {
    to int
    cost int
}

func solution(n int, edge [][]int) int {
    adj := [20001][]*Edge{}
    costs := [20001]int{}

    for i:=2; i<=n; i++ {
        costs[i] = math.MaxInt64
    }

    for _, e := range edge {
        from := e[0]
        to := e[1]
        adj[from] = append(adj[from], &Edge{from: from, to: to, weight: 1})
        adj[to] = append(adj[to], &Edge{from: to, to: from, weight: 1})
    }

    h := &Heap{Cost{to:1, cost:0}}
    heap.Init(h)

    for (*h).Len() != 0{
        e := h.Pop().(Cost)
        if costs[e.to] < e.cost {
            continue
        }

        for _, v := range adj[e.to] {            
            if costs[v.to] > costs[e.to] + v.weight {
                costs[v.to] = costs[e.to] + v.weight
                h.Push(Cost{to: v.to, cost: costs[v.to]})
            }
        }
    }

    max := 0
    maxCount := 0
    for _, v := range costs {
        if v > max {
            max = v            
            maxCount = 1
        } else if v == max {
            maxCount++
        }
    }

    return maxCount
}

```

위 풀이에서 우선순위 큐의 타입은 `Cost`이다.  
`Cost`는 도달 지점과 가중치로 이루어진 타입이다.  
`costs`라는 배열 변수엔 가중치를 시작 노드인 1번을 제외하고 모두 INF로 초기화 된다.  
이후 미리 `edges`를 순회하며 이웃 노드를 저장해놓은 배열을 참조하며 시작 노드에서 멀어질 수록 1씩 가중치를 증가시키며 heap을 구성해 나간다.

위 코드에서 heap은 최소 힙이며 가장 가중치가 큰 노드부터 우선 Pop 한다.