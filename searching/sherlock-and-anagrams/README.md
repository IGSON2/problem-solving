# [Sherlock and anagrams](https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?isFullScreen=true)

### 임의의 문열이 주어졌을 때, 좌우 대칭인 문자열 쌍의 수를 구하는 문제

###### 문제도 잘못 이해하고 있었음. Anagram이란, 한 단어를 구성하는 글자의 개수를 그대로 유지하면서 순서만 바꾼 단어라고 함. (ex. listen <-> silent)

---

## 나의 풀이
```go
func sherlockAndAnagrams(s string) int {
    cnt := 0
    for i := 0; i < len(s)-1; i++{
        Outer:
        for j := i; j < len(s); j++{
            left := s[i:j+1]
            if j+1 >= len(s){
                continue
            }
            right := s[i+1:j+2]
            
            for k := 0; k < len(left); k++ {
                if left[k] != right[len(right)-k-1]{
                    continue Outer
                }
            }
            cnt++
            fmt.Printf("[%s, %s]\n",left,right)
        }
    }
    return cnt
}
```
`[i:j+1]` 부분 문자열과 `[i+1:j+2]` 부분 문자열의 비교만 탐색하면 안됐다. 건너뛴 탐색이 너무 많았다.

문자열을 이중으로 돌며 만들어지는 모든 조합을 정렬한 뒤, 중복을 체크하면 간단히 해결 가능하다.

문자열을 정렬하여 키로 체크하는 것과, 같은 문자열을 조합하는 공식이 이해가 잘 안간다.

---
## 모범 풀이
```go
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	var cases int
	fmt.Scanf("%d", &cases)

	for t := 0; t < cases; t++ {
		var s string
		fmt.Scanf("%s", &s)

		hashtable := make(map[string]int)
		var counter int
		for length := 1; length < len(s); length++ {
			for i := 0; i+length <= len(s); i++ {
				first := i
				last := i + length
				hash := SortString(s[first:last])
				_, ok := hashtable[hash]
				if !ok {
					hashtable[hash] = 1
				} else {
					hashtable[hash]++
				}
			}
		}
		for _, v := range hashtable {
			if v > 1 {
				counter += (v - 1) * v / 2
			}
		}
		fmt.Println(counter)
	}
}
```