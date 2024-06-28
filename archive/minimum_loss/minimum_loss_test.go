package minimumloss

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumLoss(t *testing.T) {
	data, err := os.ReadFile("case2")
	require.NoError(t, err)

	dataArr := strings.Split(string(data), " ")

	testData := make([]int64, 0)

	for _, v := range dataArr {
		i, err := strconv.Atoi(v)
		require.NoError(t, err)
		testData = append(testData, int64(i))
	}

	result := minimumLoss(testData)
	t.Log("result : ", result)
}
