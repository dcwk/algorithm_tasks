package task5

func LongestCommonPrefix(strs []string) string {
	result := ""
	currentResult := ""
	countWords := len(strs)
	if countWords == 0 {
		return result
	}

	currentLetter := ""
	isNeedStop := false
	for i := 0; i < len(strs[0]); i++ {
		currentLetter = string(strs[0][i])
		for j := 1; j < countWords; j++ {
			if isNeedStop == true {
				continue
			}

			if len(strs[j]) <= i {
				isNeedStop = true
				continue
			}

			if string(strs[j][i]) != currentLetter {
				isNeedStop = true
			}
		}

		if isNeedStop == false {
			currentResult += currentLetter
		}

		if len(currentResult) > len(result) {
			result = currentResult
		}
	}

	return result
}
