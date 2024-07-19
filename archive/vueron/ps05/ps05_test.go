package ps05

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPs05(t *testing.T) {
	f, err := os.Open("testcase")
	require.NoError(t, err)
	defer f.Close()

	var N int
	_, err = fmt.Fscan(f, &N)
	require.NoError(t, err)

	arr := make([][]uint8, N)
	for i := 0; i < N; i++ {
		var line string
		_, err = fmt.Fscan(f, &line)
		require.NoError(t, err)
		cellStrs := strings.Split(line, "")
		for _, c := range cellStrs {
			if c == "0" {
				arr[i] = append(arr[i], 0)
			} else {
				arr[i] = append(arr[i], 1)
			}
		}
	}

	ps05(N, arr)

}
