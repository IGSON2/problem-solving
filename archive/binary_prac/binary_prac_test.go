package binaryprac

import "testing"

func TestBinaryPrac(t *testing.T) {
	testArray := []int32{100, 90, 90, 80, 75, 60}
	result := binarySearch(len(testArray)/2, testArray, 100, 50)
	t.Log(result)
}
