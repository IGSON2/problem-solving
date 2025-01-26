# [ 단어 변환 ](https://school.programmers.co.kr/learn/courses/30/lessons/43163?language=go)

## 모범 풀이
dfs/bsf 를 활용해 시작단어에서 목표단어까지 몇번의 변환이 필요한지 찾는 문제

맞게 풀었는데 왜 하나가 틀리는지 모르겠음

최단거리 탐색은 bsf를 활용하고 depth 또한 추적해야함.
dfs 또한 cnt를 파라미터로 하여 depth를 추적할 수 있음.

정답자 풀이
```go
func solution(begin string, target string, words []string) int {
    includeTarget := checkTarget(target, words)
    cnt := 0
    if includeTarget {
        cnt = bfs(begin, target, words)
    }
    return cnt
}

type node struct {
    word string
    level int
}

func bfs(begin string, target string, words []string) int {
    queue := []node{}
    visited := make([]bool, len(words))
    first := node{
        word: begin,
        level: 1,
    }
    queue = append(queue, first) 
    for len(queue) > 0 {
        pi := queue[0]
        queue = queue[1:]
        if isTrans(pi.word, target) {
            return pi.level
        }
        for i, w := range words {
            if visited[i] == false && isTrans(pi.word, w) {
                n := node {
                    word: words[i],
                    level: pi.level + 1,
                }
                queue = append(queue, n)
                visited[i] = true
            }
        }
    }
    return 0
}

func isTrans(str1, str2 string) bool {
    diff := 0
    for i := range str1 {
        if str1[i] != str2[i] {
            diff++
        }
    }
    return diff == 1
}

func checkTarget(target string, words []string) bool {
    for _, str := range words {
        if target == str {
            return true
        }
    }
    return false
}
```

DFS를 이용한 풀이
``` go
var answer int
func search(begin string, target string, words []string, visit []bool,cnt int){
    if cnt > len(words){
        return
    }
    if begin == target{
        if answer > cnt{
            answer = cnt
        }
        return
    }
    for idx, word := range(words){
        count := 0
        for i:=0; i<len(word); i++{
            if begin[i] != word[i]{
                count++
            }
        }
        if count == 1 && !visit[idx] {
            visit[idx] = true
            search(word,target,words,visit,cnt+1)
            visit[idx] = false
        }
    }
}
func solution(begin string, target string, words []string) int {
    answer = len(words) + 1
    visit := make([]bool, len(words))
    search(begin,target,words,visit,0)
    if answer == len(words) + 1{
        answer = 0
    }
    return answer
}
```