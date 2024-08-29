package binaryconversion

import (
	"strconv"
	"strings"
)

func solution(s string) []int {
	var cnt, removed0 int
	for ; len(s) > 1; cnt++ {

		replaced := strings.ReplaceAll(s, "0", "")
		removed0 += len(s) - len(replaced)
		s = convertBinary(len(replaced))
	}
	return []int{cnt, removed0}
}

func convertBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
