package task1

import "testing"

func TestCanCheckNonUniqueString(t *testing.T) {
	isUnique := IsUniqueChars("ghsjkahgjashgjkhasdjklghh")

	if isUnique {
		t.Error("String is not unique")
	}
}

func TestCanCheckUniqueString(t *testing.T) {
	isUnique := IsUniqueChars("qwerty")

	if !isUnique {
		t.Error("String is unique")
	}
}
