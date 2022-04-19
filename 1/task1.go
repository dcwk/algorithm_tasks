package task1

func IsUniqueChars(input string) bool {
	symbols := [128]bool{}

	for _, char := range input {
		if !symbols[char] {
			symbols[char] = true
		} else {
			return false
		}
	}

	return true
}
