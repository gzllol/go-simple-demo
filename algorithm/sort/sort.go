package sort

import "fmt"

/*
 */
func InsertionSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	// l := len(arr)
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func PrintArray(arr []int) {
	for s := range arr {
		fmt.Printf("%d ", s)
	}
}
