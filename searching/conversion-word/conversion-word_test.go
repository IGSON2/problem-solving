package conversionword

import "testing"

func TestSolution(t *testing.T) {
	cnt := solution("aaaabb", "abbbbb", []string{"aaaabb", "aaabbb", "aabbbb", "abbbbb"})
	t.Log(cnt)
}
