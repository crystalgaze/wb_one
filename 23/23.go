package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	indexToDelete := 1
	copy(slice[indexToDelete:], slice[indexToDelete+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
