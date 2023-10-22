package leetcode

func RemoveDuplicates(nums []int) int {
	var (
		unique  int
		prevNum int
	)

	for i := 0; i < len(nums); i++ {
		curNum := nums[i]

		if i > 0 && prevNum == curNum {
			if i < len(nums)-1 {
				nums = append(nums[:i-1], nums[i:]...)
				i--
			}

			continue
		}

		prevNum = curNum
		unique++
	}

	return unique
}

func RemoveMoreThanTwoDuplicates(nums []int) int {
	var (
		unique       int
		prevNum      int
		prevNumCount int
	)

	for i := 0; i < len(nums); i++ {
		curNum := nums[i]

		if i > 0 && prevNum == curNum {
			if prevNumCount > 1 {
				nums = append(nums[:i-1], nums[i:]...)
				i--
			}
			prevNumCount += 1

			continue
		}

		prevNum = curNum
		prevNumCount = 1
		unique++
	}

	return len(nums)
}
