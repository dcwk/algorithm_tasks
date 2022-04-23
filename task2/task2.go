package task2

func IsPermutationOfPalindrome(input string) bool {
	hashMap := make(map[string]int)
	palindromeSeparator := ""

	for _, char := range input {
		hashMap[string(char)]++
	}

	for charValue, charCount := range hashMap {
		if charCount%2 == 1 && palindromeSeparator != "" {
			return false
		}

		if charCount == 1 && palindromeSeparator == "" {
			palindromeSeparator = string(charValue)
		} else if charCount == 1 && palindromeSeparator != "" {
			return false
		}
	}

	return true
}
