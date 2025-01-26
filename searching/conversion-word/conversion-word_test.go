package conversionword

import "testing"

func TestSolution(t *testing.T) {
	cnt := solution("cot", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	t.Log(cnt)
}
