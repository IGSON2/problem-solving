package gridlandmetro

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGridlandMetro(t *testing.T) {
	data, err := os.ReadFile("case1")
	require.NoError(t, err)

	dataArr := strings.Split(string(data), "\n")

	firstLint := strings.Split(dataArr[0], " ")

	n, err := strconv.Atoi(firstLint[0])
	require.NoError(t, err)
	m, err := strconv.Atoi(firstLint[1])
	require.NoError(t, err)
	k, err := strconv.Atoi(firstLint[2])
	require.NoError(t, err)

	tracks := make([][]int32, 0)

	for i := 1; i <= k; i++ {
		track := strings.Split(dataArr[i], " ")
		row, err := strconv.Atoi(track[0])
		require.NoError(t, err)
		start, err := strconv.Atoi(track[1])
		require.NoError(t, err)
		end, err := strconv.Atoi(track[2])
		require.NoError(t, err)
		tracks = append(tracks, []int32{int32(row), int32(start), int32(end)})
	}
	result := gridlandMetro(int32(n), int32(m), int32(k), tracks)
	t.Log("result : ", result)
}

func TestGridlandMetro2(t *testing.T) {
	n, m := 10, 10
	tracks := [][]int32{
		{1, 1, 8},
		{1, 1, 7},   // [1,8] => 8
		{1, 7, 9},   // [1,9] => 9
		{1, 10, 10}, // [1,9] + [10,10] => 9 + 1 = 10
		{2, 8, 8},   //
		{2, 3, 10},  //
		{2, 3, 3},   // [4,10] + [3,3] => 7 + 1 = 8
	}
	expected := uint64(82)
	result := gridlandMetro(int32(n), int32(m), int32(len(tracks)), tracks) // expected 11
	require.Equal(t, expected, result)
}
