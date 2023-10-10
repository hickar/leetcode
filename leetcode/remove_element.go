package leetcode

func RemoveElement(nums []int, val int) int {
	var notEqualCount int

	for i := 0; i < len(nums); i++ {
		curItem := nums[i]
		if curItem != val {
			notEqualCount++
		} else {
			if i < len(nums)-1 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}

	return notEqualCount
}
