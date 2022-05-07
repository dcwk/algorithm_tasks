package task6

import (
	"fmt"
	"testing"
)

func TestCanConvertIntToRoman(t *testing.T) {
	result := IntToRoman(1994)
	if result != "MCMXCIV" {
		t.Error(fmt.Sprintf("Expected MCMXCIV result: %s", result))
	}
}
