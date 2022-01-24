package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Accum("Helo"))
}

func Accum(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		letter := string(s[i])
		for j := 0; j < i+1; j++ {
			if j == 0 {
				letter = strings.ToUpper(letter)
			} else {
				letter = strings.ToLower(letter)
			}
			result += letter
		}
		result += "-"
	}
	result = strings.TrimSuffix(result, "-")
	return result
}
