# [Red knights shortest path](https://www.hackerrank.com/challenges/red-knights-shortest-path/problem)

### 이차원 공간에서 총 6방향으로 이동할 수 있는 Red knight가 가장 빠르게 이동할 수 있는 경로를 탐색하는 문제. 방향의 우선순위가 있어 같은 길이여도 우선적인 방향으로 진행한 경로가 정답이 된다.

---

## 나의 풀이

```go

type MovingPoint struct{
    i,j int
    DestiFromCell string
}

type PointsByCell struct{
    Mp [][][]MovingPoint
    Visied [][]bool
    Results map[int][]string
}

var directions = [6]MovingPoint{
    {-2,-1,"UL"},
    {-2,1,"UR"},
    {0,2,"R"},
    {2,1,"LR"},
    {2,-1,"LL"},
    {0,-2,"L"},
}

func NewPointsByCell(n int) *PointsByCell{
    p := new(PointsByCell)
    temp := make([][][]MovingPoint,n)
    vTemp := make([][]bool,n)
    for i := 0; i < n; i++{
        temp[i] = make([][]MovingPoint,n)
        vTemp[i] = make([]bool,n)
        for j := 0; j < n; j++{
            for _,d := range directions{
                newI,newJ := i+d.i, j+d.j
                if newI < 0 || newJ < 0 || newI > len(temp)-1 || newJ > len(temp[0])-1{
                    continue
                }
                point := MovingPoint{
                        newI,newJ,d.DestiFromCell,
                        }
                temp[i][j]=append(temp[i][j], point)
            }
        }
    }
    p.Mp = temp
    p.Visied = vTemp
    p.Results = make(map[int][]string)
    return p
}

func (p PointsByCell)findPoints(i,j int) []MovingPoint{
    return p.Mp[i][j]
}

func dfs(startI, startJ, endI, endJ int, chessMap *PointsByCell, stack []string) bool {
    i,j := startI,startJ
        
    if chessMap.Visied[i][j]{
        return false
    }
    
    chessMap.Visied[i][j] = true
    
    // 잘못된 경로를 탐색하며 visited에 체크된 Cell이 정답 경로중 한 Cell일 때 올바른 처리 불가
    for _, mp := range chessMap.Mp[i][j]{
        if mp.i == endI && mp.j == endJ{
            tempResult := make([]string,len(stack)+1)
            copy(tempResult,append(stack,mp.DestiFromCell))
            if _,exist := chessMap.Results[len(tempResult)]; !exist{
                chessMap.Results[len(tempResult)] = tempResult
            }
            return false
        }
    }
    
    var isReachable bool
    for _, mp := range chessMap.Mp[i][j]{
        stack = append(stack, mp.DestiFromCell)
        isReachable := dfs(mp.i,mp.j,endI,endJ,chessMap,stack)
        if !isReachable{
            stack = stack[:len(stack)-1]
        }
    }
    
    if !isReachable{
        return false
    }
    
    return true
}

func printShortestPath(n int32, i_start int32, j_start int32, i_end int32, j_end int32) {
    chessMap := NewPointsByCell(int(n))
    
    dfs(int(i_start),int(j_start),int(i_end),int(j_end),chessMap,[]string{})
    
    
    minPath := int(n*n*8)
    result := []string{}
    
    for k,v := range chessMap.Results{
        if k < minPath{
            minPath = k
            result = v
        }
    }
    
    if len(result) > 0 {
        fmt.Println(len(result))
        fmt.Println(strings.Join(result," "))
        return
    }
    
    fmt.Println("Impossible")

}
```

12개의 케이스 중 7개의 Case에서 Wrong answer에 걸렸다.
재귀호출로 인해 잘못된 경로 탐색 중 방문처리 했던 셀이 정답 경로중 한 셀일 때 올바른 처리를 하지 못했다.

---

## 모범 풀이
```go
type point struct {
    x int
    y int
    p *point
    steps int
    stepType string
}

func (p *point) isOutside(n int) bool {
    return p.x < 0 || p.x >= n || p.y < 0 || p.y >= n
}

var steps = []struct{
    dx int
    dy int
    stepType string
}{
    {stepType: "UL", dx: -2, dy: -1},
    {stepType: "UR", dx: -2, dy: 1},
    {stepType: "R", dx: 0, dy: 2},
    {stepType: "LR", dx: 2, dy: 1},
    {stepType: "LL", dx: 2, dy: -1},
    {stepType: "L", dx: 0, dy: -2},
}

func main() {
    var n int
    fmt.Scan(&n)
    var visited [][]bool
    visited = make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, n)
    }

    var start point
    fmt.Scan(&start.x, &start.y)
    var end point
    fmt.Scan(&end.x, &end.y)

    queue := []*point{&start}
    visited[start.x][start.y] = true
    
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        
        if p.x == end.x && p.y == end.y {
            printFound(p)
            return
        }
        
        for _, step := range steps {
            np := &point{
                x: p.x + step.dx,
                y: p.y + step.dy,
                p: p,
                steps: p.steps + 1,
                stepType: step.stepType,
            }
            if np.isOutside(n) {
                continue
            }
            if visited[np.x][np.y] {
                continue
            }
            visited[np.x][np.y] = true
            queue = append(queue, np)
        }
        
    }
    
    fmt.Println("Impossible")
}

func printFound(p *point) {
    fmt.Println(p.steps)
    drs := make([]string, 0, p.steps)
    for p.p != nil {
        drs = append(drs, p.stepType)
        p = p.p
    }
    
    for i := len(drs) - 1; i >= 0; i-- {
        fmt.Print(drs[i])
        fmt.Print(" ")
    }
    fmt.Println()
}
```

링크드 리스트를 활용하여 이전 Cell을 기록하였다. 그런데 의문인점은 최단경로가 아닌 경로로 도작점에 도달해도 통과하지 않느냐는 것이다.

```go
type pos struct {
    y, x int
}

func makemove(p pos, name string) pos {
    var n pos
    
    switch name {
        case "UL": n = pos{p.y-2, p.x-1}
        case "UR": n = pos{p.y-2, p.x+1}
        case "R":  n = pos{p.y, p.x+2}
        case "LR": n = pos{p.y+2, p.x+1}
        case "LL": n = pos{p.y+2, p.x-1}
        case "L":  n = pos{p.y, p.x-2}                       
        
    }
    
    return n
}

type move struct {
    pos pos
    name string
}

func possible(n int, start pos, end pos) []string {
    visited := map[pos]move{}
    queue := []pos{start}
    moves := []string{"UL", "UR", "R", "LR", "LL", "L"}
    
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        
        if p == end {
            break
        }
        
        for _, name := range moves {
            pn := makemove(p, name)
            if pn.x < 0 || pn.x >= n || pn.y < 0 || pn.y >= n {
                continue
            }
            
            if _, ok := visited[pn]; !ok {
                visited[pn] = move{p, name}
                queue = append(queue, pn)
            }
        }           
    }
    
    if _, ok := visited[end]; !ok {
        return nil
    }
    
    l := []string{}
    for p := end; p != start; p = visited[p].pos {        
        l = append(l, visited[p].name)
    }
    return l
}
```
BFS를 사용하여 풀이하였다.
보면 인접셀들을 굳이 기록하지 않는다. 어차피 범위 내에서 지정한 방향과 칸 수 만큼 움직이다 도착점을 찾으면 되기 때문, 추가로 moves 배열엔 우선순위대로 이동 방향이 삽입되어 있다.
queue의 길이가 1일 때 queue[1:]은 옳은 표현이다. 인덱스가 슬라이스의 길이와 같으면 빈 슬라이스를 반환한다.