package sort

func merge(left, right []int) []int {
	result := make([]int, len(left) + len(right))
	i, j, k := 0, 0, 0

	for ; i < len(left) && j < len(right); k++ {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	for ; i < len(left); k++ {
		result[k] = left[i]
		i++
	}

	for ; j < len(right); k++ {
		result[k] = right[j]
		j++
	}

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

//func MergeInsertionSort(arr []int) []int {
//	if len(arr) > threshold {
//		if len(arr) == 1 {
//			return arr
//		}
//
//		left := MergeInsertionSort(arr)
//		right := MergeInsertionSort(arr)
//		merge(left, right)
//	} else {
//		return func(arr []int) []int {
//			for j := 0; j < len(arr); j++ {
//				key := arr[j]
//				i := j - 1
//
//				for i >= 0 && arr[i] > key {
//					arr[i+1] = arr[i]
//					i--
//				}
//
//				arr[i+1] = key
//			}
//		}(arr)
//	}
//}