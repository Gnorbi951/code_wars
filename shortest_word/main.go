package main

import (
	"strings"
)

func main() {
	FindShort("bitcoin take over the world maybe who knows perhaps")
	FindShort("turns out random test cases are easier than writing out basic ones")
}

func FindShort(s string) int {
	arr := strings.Split(s, " ")

	shortest := 100

	for _, element := range arr {
		if len(element) < shortest {
			shortest = len(element)
		}
	}
	return shortest
}
