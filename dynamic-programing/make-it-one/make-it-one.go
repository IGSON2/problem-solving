package makeitone

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// n까지의 최소 연산을 배열에 상향식으로 기록해 나가야 함.
// DP문제로 보조 자료구조가 필요했음.
func makeItOne(n int) {

	arr := make([]int, n+1)
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + 1

		if i%3 == 0 {
			arr[i] = min(arr[i/3]+1, arr[i])
		}

		if i%2 == 0 {
			arr[i] = min(arr[i/2]+1, arr[i])
		}
	}

	fmt.Println(arr[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // 원래 나의 풀이 - 하향식 진행
// func makeItOne(n int) {

// 	result := 0

// 	for n > 1 {
// 		fmt.Println(n)
// 		if n%3 == 0 {
// 			n /= 3
// 		} else if n%3 == 1 {
// 			n--
// 		} else {
// 			if n%2 == 0 {
// 				n /= 2
// 			} else {
// 				n--
// 			}
// 		}
// 		result++
// 	}

// 	fmt.Println(result)
// }
