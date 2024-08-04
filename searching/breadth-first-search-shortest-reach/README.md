# [Really special subtree](https://www.hackerrank.com/challenges/bfsshortreach/problem?isFullScreen=true)

### 그래프에서 임의의 시작점을 기준으로 모든 노드까지 각각의 최소 거리를 구하는 문제

---

## 나의 풀이

```go
func bfs(n int32, m int32, G map[int][]int, s int) []int32 {
    result := make([]int32,n+1)
    visited := make([]bool,n+1)
    
    queue := make([]int,1)
    queue[0] = s
    
    for distance := int32(6); len(queue) > 0; {
        v := int(queue[0])
        queue = queue[1:]
        
        visited[v] = true
        
        var isReachable bool
        var prevNode int
        for _,node := range G[v]{
            if !visited[node] && node != prevNode{
                result[node] = distance
                queue = append(queue, node)
                isReachable = true
                prevNode = node
            }
        }
        // fmt.Printf("node : %d\t%v\n",v,queue)
        if isReachable {
            distance += 6
        }
    }
    
    for i,isVisited := range visited{
        if !isVisited{
            result[i] = -1
        }
    }
    
    return result[1:]
}
```

시작 노드 방문 체크를 따로 실시하고 방문 노드 체크를 큐에 추가할 때 해 주어야 한다.
그리고 distance를 따로 증가시킬 필요가 없으며, 인접노드의 거리 + 6만 해주면 된다.

---

## 수정된 풀이

```go
visited[s] = true
    
    for len(queue) > 0 {
        v := int(queue[0])
        queue = queue[1:]
        
        
        var prevNode int
        for _,node := range G[v]{
            if !visited[node] && node != prevNode{
                result[node] = result[v] + 6 // 인접노드의 거리 +6
                queue = append(queue, node)
                visited[node] = true // queue에 추가하기 전 visited 체크. 그래야 다른 노드에 공통 접접이 있을 경우 알고리즘이 꼬이지 않음.
                prevNode = node
            }
        }
    }
```

---

## dijkstra를 이용한 풀이

```go
var INF = math.MaxInt

type Item struct {
    value, priority, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue)Len()int{
    return len(pq)
}

func (pq PriorityQueue)Less(i,j int)bool{
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue)Swap(i,j int){
    pq[i],pq[j] = pq[j],pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue)Push(x interface{}){
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue)Pop() interface{}{
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1
    *pq = old[:n-1]
    return item
}

func (pq *PriorityQueue)update(item *Item, value, priority int){
    item.value = value
    item.priority = priority
    heap.Fix(pq,item.index)
}

func dijkstra(G []*list.List, s,n int)[]int{
    pq := make(PriorityQueue,G[s].Len())
    dists := make([]int,n+1)
    for ix := 1; ix < n+1; ix++ { // 1번 노드부터 마지막 노드까지 거리를 INF로 초기화
        dists[ix] = INF
    }
    
    i,c := 0,0
    for e := G[s].Front(); e != nil; e = e.Next(){ 
        c = e.Value.(int)
        pq[i] = &Item{value: c,priority:1, index:i} // 시작점의 이웃 노드를 우선순위 큐에 추가
        i++
        dists[c] = 1 // 시작점의 이웃 노드의 거리를 1로 변경
    }
    
    heap.Init(&pq) // pq를 heap 자료형에 맞게 설정
    
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        
        for adj := G[item.value].Front(); adj != nil; adj = adj.Next(){
            u := adj.Value.(int)
            if u == s{
                continue
            }
            if dists[item.value] != INF && dists[item.value] + 1 < dists[u]{  // 이미 n 번 째 노드까지 도달했고 기존 n의 인접노드 u까지의 경로보다 짧은 경로를 찾았을 때
                dists[u] = dists[item.value] + 1
                im := &Item{
                    value : u,
                    priority: dists[u],
                }
                heap.Push(&pq,im)
                pq.update(im,im.value,dists[u])
            }
        }
    }
    return dists
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int(nTemp)

        mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        m := int32(mTemp)

        G := make([]*list.List,n+1)
        for i := 0; i < int(nTemp+1); i++{
            G[i] = list.New()
        }
        
        for i := 0; i < int(m); i++ {
            edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

            from, _ := strconv.Atoi(edgesRowTemp[0])
            to, _ := strconv.Atoi(edgesRowTemp[1])
            G[from].PushBack(to)
            G[to].PushBack(from)
        }

        sTemp, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
        checkError(err)
        s := sTemp

        result := dijkstra(G,s,n)
        
        for ix := 1; ix < n+1; ix++ {
            if ix != s {
                if result[ix] != INF {
                    fmt.Fprintf(writer,"%d", result[ix]*6)
                } else {
                    fmt.Fprintf(writer,"%d", -1)
                }
                if ix < n {
                    fmt.Fprint(writer," ")
                }
            }
        }

        fmt.Fprintf(writer, "\n")
    }

    writer.Flush()
}
```