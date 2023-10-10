package leetcode

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i, num := range nums2 {
		nums1[m+i] = num
	}

	for i := 1; i < len(nums1); i++ {
		key := nums1[i]
		j := i - 1

		for j >= 0 && nums1[j] > key {
			nums1[j+1] = nums1[j]
			j--
		}

		nums1[j+1] = key
	}
}
