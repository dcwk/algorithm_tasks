package task4

func RomanToInt(s string) int {
	result := 0
	if len(s) == 0 {
		return result
	}

	alphabet := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	letter := ""
	nextChar := ""

	for i := 0; i < len(s); i++ {
		nextChar = ""
		letter = string(s[i])
		if i != len(s)-1 {
			nextChar = string(s[i+1])
		}

		if alphabet[letter+nextChar] != 0 {
			result += alphabet[letter+nextChar]
			i++
			continue
		}

		result += alphabet[letter]
	}

	return result
}
