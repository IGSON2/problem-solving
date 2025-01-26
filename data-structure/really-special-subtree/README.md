# [Really special subtree](https://www.hackerrank.com/challenges/kruskalmstrsub/problem?isFullScreen=true)

### 그래프에서 모든 노드로 구성된 하위 그래프의 최소 가중치를 구하는 문제

---

## 나의 풀이

```go
    for i := 0; i < int(gEdges); i++ {
        edgeFromToWeight := strings.Split(readLine(reader), " ")

        edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
        checkError(err)

        edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
        checkError(err)

        edgeWeight, err := strconv.ParseInt(edgeFromToWeight[2], 10, 64)
        checkError(err)

        edges[int32(edgeFrom)] = append(edges[int32(edgeFrom)],Weight{
            to: int32(edgeTo),
            weight:int32(edgeWeight),
            })
        edges[int32(edgeTo)] = append(edges[int32(edgeTo)],Weight{
            to: int32(edgeFrom),
            weight:int32(edgeWeight),
            })
    }
    
    
    for _,v := range edges{
        sort.Slice(v,func(i,j int)bool{
            return v[i].weight < v[j].weight
        })
    }
    
    weight := int32(0)
    check := make(map[edge]bool)
    
    for i := 1; i < int(gNodes); i++{
        for _, e := range edges[int32(i)]{
            if !check[edge{from:int32(i),to:e.to}]{
                weight += e.weight
                check[edge{from:e.to,to:int32(i)}] = true
                fmt.Println(i,e)
                break
            }
        }
    }
    
    fmt.Fprintln(writer,weight)
```

마지막 노드가 이미 포함되어있을 수 있고, 아닐수도 있는 경우를 일일이 고려해야 한다는 점에서 알고리즘을 잘못 짰다고 느낌.

---
## 생소한 개념

#### 우선순위 큐 (Priory queue)
우선순위의 개념을 큐에 도입한 자료 구조이다. 데이터들이 **우선순위**를 가지고 있고 우선순위가 제일 높은 데이터를 POP한다.  
시뮬레이션 시스템, 네트워크 트래픽 제어, 운영체제의 작업 스케쥴링 등에 사용되는 자료구조이다.  
우선순위 큐는 배열, 연결리스트, 힙 으로 구현이 가능하다. 이 중에서 힙(heap)으로 구현하는 것이 가장 효율적이다

#### 힙 (heap) 자료구조
완전 이진 트리의 일종으로 우선순위 큐를 위하여 만들어진 자료구조이다.  
여러 개의 값들 중에서 최댓값이나 최솟값을 빠르게 찾아내도록 만들어진 자료구조이다.  
힙은 일종의 반정렬 상태(느슨한 정렬 상태) 를 유지한다.  
최대 힙과 최소 힙으로 나뉘어진다. 최대 힙에선 부모 노드가 자식 노드보다 같거나 크고, 최소 힙에선 부모 노드가 자식 노드보다 같거나 작은 구조를 가진다.


<img src="https://media.geeksforgeeks.org/wp-content/cdn-uploads/20221220165711/MinHeapAndMaxHeap1.png" />


힙을 저장하는 표준적인 자료구조는 배열이며, 구현을 쉽게 하기 위하여 배열의 첫 번째 인덱스인 0은 사용되지 않는다. 또한, 특정 위치의 노드 번호는 새로운 노드가 추가되어도 변하지 않는다.  
##### 힙에서 부모 노드와 자식 노드의 관계
```
왼쪽 자식의 인덱스 = (부모의 인덱스) * 2
오른쪽 자식의 인덱스 = (부모의 인덱스) * 2 + 1
부모의 인덱스 = (자식의 인덱스) / 2
```

#### 최소 스패닝 트리 (MST)
스패닝 트리(신장 트리, Spanning Tree)란, **그래프의 모든 정점을 잇지만 사이클이 없는 부분 그래프를 의미**한다.  

이 스패닝 트리 중 간선의 가중치 합이 가장 적은 것이 최소 스패닝 트리이다.  

스패닝 트리 내 V개의 모든 정점을 연결하는 간선의 수는 V - 1개이다.  

<img src="https://media.geeksforgeeks.org/wp-content/uploads/20200316173940/Untitled-Diagram66-3.jpg"/>


#### 다익스트라 알고리즘
다이나믹 프로그래밍을 활용한 대표적인 최단 경로 탐색 알고리즘이다. 다익스트라 알고리즘을 사용하면 특정한 하나의 정점에서 나머지 모든 정점으로 가는 최단 경로를 알아낼 수 있다.  

흔히 인공위성 GPS를 이용한 소프트웨어에서 많이 사용하며, 


## 모범 풀이

```go
type Item struct {
    value    int
    priority int
    index int
}

type PriorityQueue []*Item

func (pq PriorityQueue)Len()int{
    return len(pq)
}

func (pq PriorityQueue)Less(i,j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue)Swap(i,j int) {
    pq[i],pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

// Push와 Pop은 힙 구조를 직접 변경하기 때문에 타입의 값이 아닌 포인터로 정의해야 한다.
func (pq *PriorityQueue)Push (x interface{}){
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue)Pop() interface{}{
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[0:n-1]
    return item
}

// pq내 아이템의 우선순위와 값을 변경한다.
func (pq *PriorityQueue)update(item *Item, value int, priority int){
    item.value = value
    item.priority = priority
    heap.Fix(pq,item.index)
}
```