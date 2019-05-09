package binarysearch

func BinarySearch(arr []int, a int) int {
	s := 0
	e := len(arr) - 1
	var mid int
	for s <= e {
		mid = (s + e) / 2
		if arr[mid] == a {
			return mid
		} else if arr[mid] < a {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return -1
}
