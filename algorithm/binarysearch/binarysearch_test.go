package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 4, 6}
	if 2 != BinarySearch(arr, 4) {
		t.Error("failed")
	}
}
