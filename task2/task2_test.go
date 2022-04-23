package task2

import "testing"

func TestCanCheckStringIsNotPalindrom(t *testing.T) {
	isPalindrome := IsPermutationOfPalindrome("jlmmkljkhkl")
	if isPalindrome == true {
		t.Error("String is not palindrome")
	}
}

func TestCanCheckStringIsPalindrom(t *testing.T) {
	isPalindrome := IsPermutationOfPalindrome("tactcoa")
	if isPalindrome == false {
		t.Error("String is palindrome")
	}
}
