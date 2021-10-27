package main

import "strings"

// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.

var vowels = []string{"a", "e", "i", "o", "u"}

func GetCount(str string) (count int) {
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if contains(vowels, char) {
			count++
		}
	}
	return count
}

func contains(arr []string, str string) bool {
	for _, e := range arr {
		if e == strings.ToLower(str) {
			return true
		}
	}
	return false
}

func main() {
	GetCount("Azbeszt megb*szta")
	GetCount("Lajos")
}
