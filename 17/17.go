package main

import (
	"fmt"
)

func binarySearch(x int, input []int) int {
	low := 0
	high := len(input) - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if input[mid] == x {
			return mid
		} else if input[mid] < x {
			low = mid + 1
		} else if input[mid] > x {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	data := []int{1, 10, 27, 45, 69, 71, 74, 75, 81, 95, 101, 141}
	x := 45
	fmt.Println("Индекс X в data равен", binarySearch(x, data))
}
