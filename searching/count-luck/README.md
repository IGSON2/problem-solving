# [Count luck](https://www.hackerrank.com/challenges/count-luck/copy-from/389338286?isFullScreen=true)

### 이중 배열에서 목적지를 찾아가는 도중 갈림길이 있을 때, 갈림길의 수를 맞추는 문제이다.

---

## 나의 풀이
```go
var directions = [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}

func dfs(matrix [][]string, visited [][]bool, i,j int, cnt *int) bool{
    if 
    i < 0 || 
    j < 0 || 
    i > len(matrix)-1 || 
    j > len(matrix[0])-1||
    matrix[i][j] == "X" ||
    visited[i][j] {
        return false
    }
    
    if matrix[i][j] == "*"{
        return true
    }
    
    visited[i][j] = true
    
    var ways int
    
    for _,d := range directions {
        newI, newJ := i + d[0], j + d[1]
        isReachable := dfs(matrix,visited,newI,newJ,cnt)
        if isReachable{
            ways++
        }
    }
    
    if ways >= 2{
        *cnt++
    }
    
    return true
}

func countLuck(matrix [][]string, k int32, i,j int) string {
    var cnt int
    visited := make([][]bool,len(matrix))
    
    for i,_ := range visited{
        visited[i] = make([]bool, len(matrix[0])) 
    }
    
    dfs(matrix,visited,i,j,&cnt)
    
    result := "Oops!"
    if cnt == int(k){
        result = "Impressed"
    }
    return result
}
```

8개의 케이스 중 4개의 Case에서 Wrong answer에 걸렸다.
재귀호출로 인해 목적지로의 경로가 아닌 잘못된 길의 교차로까지 카운팅하는 게 문제였다.

---

## 모범 풀이
```go
func dfs(node, end, choices, k int, seen *map[int]bool, graph *[][]int) {
	(*seen)[node] = true
	unseen := 0
	for _, v := range (*graph)[node] {
		if !(*seen)[v] {
			unseen++
		}
	}
	if unseen > 1 {
		choices++
	}
	for _, next := range (*graph)[node] {
		if next == end {
			if k == choices {
				fmt.Println("Impressed") // 목적지에 도달하기 전, 다른 갈랫길에서 choices가 가산된 상태라면?
			} else {
				fmt.Println("Oops!")
			}
			return
		}
		if !(*seen)[next] {
			dfs(next, end, choices, k, seen, graph)
		}
	}
}

func main() {
	var cases int
	fmt.Scanln(&cases)

	for t := 0; t < cases; t++ {
		var n, m, k int
		fmt.Scanln(&n, &m)

		forest := make([]string, n) // 이차원 배열이 아닌 일차원 배열로 초기화
		var exit, hermione int
		adjacencyList := make([][]int, n*m)
		for i := 0; i < n; i++ {
			fmt.Scanln(&forest[i])
			isLocation := strings.Index(forest[i], "*") // 종료셀
			if isLocation > -1 {
				exit = i*m + isLocation
			}
			isLocation = strings.Index(forest[i], "M") // 시작셀
			if isLocation > -1 {
				hermione = i*m + isLocation
			}
		}
		// 인접셀을 구함
		for i := 0; i < n; i++ {
			for j, _ := range forest[i] {
				if j-1 >= 0 && forest[i][j-1] != 'X' {
					adjacencyList[i*m+j] = append(adjacencyList[i*m+j], i*m+j-1)
				}
				if j+1 < m && forest[i][j+1] != 'X' {
					adjacencyList[i*m+j] = append(adjacencyList[i*m+j], i*m+j+1)
				}
				if i-1 >= 0 && forest[i-1][j] != 'X' {
					adjacencyList[i*m+j] = append(adjacencyList[i*m+j], (i-1)*m+j)
				}
				if i+1 < n && forest[i+1][j] != 'X' {
					adjacencyList[i*m+j] = append(adjacencyList[i*m+j], (i+1)*m+j)
				}
			}
		}
		fmt.Scanln(&k)

		seen := make(map[int]bool)
		dfs(hermione, exit, 0, k, &seen, &adjacencyList)
	}
}
```

`sort.Search()`를 사용하여 풀이하였다. `sort.Search()`가 탐색시간이 더 빠른걸까?