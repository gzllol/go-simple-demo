package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	a := []int{3, 2, 1}
	InsertionSort(a)
	PrintArray(a)
}
