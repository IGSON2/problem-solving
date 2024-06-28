package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func MakeRanInt(minimum, maximum int) int {
	ranSeed := big.NewInt(int64(maximum - minimum))
	if ranSeed.Cmp(big.NewInt(0)) <= 0 {
		return 0
	}
	ranBigNum, err := rand.Int(rand.Reader, ranSeed)
	if err != nil {
		return -1
	}
	return int(ranBigNum.Int64()) + minimum
}

func MakeRanString(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		c := alphabet[MakeRanInt(0, len(alphabet))]
		sb.WriteByte(c)
	}

	return sb.String()
}
