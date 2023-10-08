package search

import "math"

func BinarySearchRecursive(arr []int, item, low, high int) int {
	mid := int(math.Ceil(float64(low) + (float64(high)-float64(low))/2))

	if high >= low {
		if arr[mid] == item {
			return mid
		}

		if item > arr[mid] {
			return BinarySearchRecursive(arr, item, mid+1, high)
		} else {
			return BinarySearchRecursive(arr, item, low, mid-1)
		}
	}

	return -1
}