package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	arr := []int{12, 331, 31, 451, 54, 52}
	fmt.Println(SortNumbers(arr))
}

func SortNumbers(numbers []int) []int {
	// My own bubblesort

	i := 1

	for true {

		if !(i < len(numbers)) {
			break
		}

		j := 0
		for j < len(numbers)-1 {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = temp
			}
			j++
		}
		i++

	}

	return numbers
}
