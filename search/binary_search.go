package search

func BinarySearch(arr []int, item int) int {
	return binSearch(arr, item, 0, len(arr)-1)
}

func binSearch(arr []int, item, low, high int) int {
	if high < low {
		return -1
	}

	mid := (low + high) / 2

	switch {
	case item > arr[mid]:
		return binSearch(arr, item, mid+1, high)
	case item < arr[mid]:
		return binSearch(arr, item, low, mid-1)
	default:
		return mid
	}
}
