package leetcode

// TODO: finish this

func SortColors(nums []int) {
	var (
		redStart  = 0
		redCursor = redStart

		blueStart  = len(nums) / 3
		blueCursor = blueStart

		greenStart  = len(nums) / 3 * 2
		greenCursor = greenStart
	)

	var (
		cursor     *int
		colorStart *int
	)
	for i := 0; i < len(nums); i++ {
		color := nums[i]

		switch color {
		case 0:
			cursor, colorStart = &redCursor, &redStart
		case 1:
			cursor, colorStart = &blueCursor, &blueStart
		case 2:
			cursor, colorStart = &greenCursor, &greenStart
		}

		if i >= *colorStart && i < *colorStart+len(nums)/3 {
			continue
		}

		nums[i], nums[*cursor] = nums[*cursor], nums[i]
		if color != getNeededColorByIndex(i, len(nums)) {
			i -= 1
		}

		*cursor += 1
	}
}

func getNeededColorByIndex(i, limit int) int {
	switch {
	case i <= limit/3:
		return 0
	case i <= limit/3*2:
		return 1
	case i <= limit:
		return 2
	}

	return 0
}
