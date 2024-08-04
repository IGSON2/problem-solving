package escapemazewithlever

import "testing"

func TestEscapeMazeWithLever(t *testing.T) {
	t.Log(solution([]string{
		"SOOOO",
		"LXXXE",
	}))
}

func TestMapAddress(t *testing.T) {
	type testStruct struct {
		i, j int
		temp int
	}
	m := make(map[testStruct][]testStruct)
	cases := []testStruct{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}}

	for _, c := range cases {
		m[c] = []testStruct{cases[0], cases[1]}
	}

	cases[0].temp = cases[1].temp + 1

	t.Log(m[cases[0]])
}
