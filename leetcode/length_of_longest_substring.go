package leetcode

// Given a string "s", find the length of the longest substring without repeating characters.

func LengthOfLongestSubstring(s string) int {
	charMap := make(map[string]bool)
	curLen := 0
	maxLen := 0

	for _, char := range s {
		if _, ok := charMap[string(char)]; !ok {
			charMap[string(char)] = true
			curLen++
			if curLen > maxLen {
				maxLen = curLen
			}
		} else {
			repeatCount := 0
			for key, _ := range charMap {
				delete(charMap, key)

				if key != string(char) {
					if repeatCount == 0 {
						curLen--
					}
				} else {
					if repeatCount == 1 {
						break
					} else {
						//curLen--
						repeatCount++
					}
				}
			}
		}
	}

	if curLen > maxLen {
		maxLen = curLen
	}

	return len(charMap)
}