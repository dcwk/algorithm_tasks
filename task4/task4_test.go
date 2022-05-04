package task4

import "testing"

func TestCanSimpleConvert(t *testing.T) {
	summ := RomanToInt("III")
	if summ != 3 {
		t.Error("Expected sum = 3!")
	}
}

func TestCanHardConvert(t *testing.T) {
	summ := RomanToInt("LVIII")
	if summ != 58 {
		t.Error("Expected sum = 58!")
	}
}

func TestCanVeryHardConvert(t *testing.T) {
	summ := RomanToInt("MCMXCIV")
	if summ != 1994 {
		t.Error("Expected sum = 1994!")
	}
}
