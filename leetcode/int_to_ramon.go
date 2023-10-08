package leetcode

func IntToRoman(num int) string {
	romanDigits := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result string

	for romanNum, romanNumVal := range romanDigits {
		switch {
		case num < romanNumVal:

		case num > romanNumVal:

		case num == romanNumVal:
			return romanNum
		}
	}

	return result
}
