package leetcode

func RomanToInteger(s string) int {
	romansMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var (
		result    int
		lastValue int
	)

	for i := len(s); i > 0; i-- {
		romanChar := s[i-1]

		currentValue := romansMap[rune(romanChar)]
		if currentValue < lastValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		lastValue = currentValue
	}

	return result
}
