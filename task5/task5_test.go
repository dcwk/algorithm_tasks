package task5

import "testing"

func TestCanGetHardPrefix(t *testing.T) {
	input := []string{"ab", "a"}

	result := LongestCommonPrefix(input)
	if result != "a" {
		t.Error("Expected result string: a")
	}
}

func TestCanSimplePrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}

	result := LongestCommonPrefix(input)
	if result != "fl" {
		t.Error("Expected result string: fl")
	}
}

func TestCanGetEmptyPrefix(t *testing.T) {
	input := []string{"dog", "racecar", "car"}

	result := LongestCommonPrefix(input)
	if result != "" {
		t.Error("Expected empty string")
	}
}
