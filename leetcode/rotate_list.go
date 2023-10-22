package leetcode

func RotateList(nums []int, k int) {
	k %= len(nums)
	if k > len(nums) {
		return
	}

	offset := len(nums) - k
	tmp := make([]int, offset)

	copy(tmp, nums[:offset])
	copy(nums[:k], nums[offset:])
	copy(nums[k:], tmp)
}
