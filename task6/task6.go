package task6

func IntToRoman(num int) string {
	result := ""
	currentNum := num
	isNeedStop := false

	if num == 0 {
		return result
	}

	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for num != 0 {
		for i := 0; i < len(values); i++ {
			if isNeedStop == true {
				continue
			}

			if values[i] > num {
				continue
			}

			currentNum = currentNum - values[i]

			if currentNum >= 0 {
				num = currentNum
				result += letters[i]
				isNeedStop = true
			}

			currentNum = num
		}

		isNeedStop = false
	}

	return result
}
