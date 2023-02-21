package main

import "fmt"

func main() {
	data1 := []int{1, 6, 435, 8579, 9, 2356, 856, 796, 2345, 8578, 5345}
	data2 := []int{9769, 56, 876, 59, 8579, 8578, 73, 1, 8217, 67, 8}
	result := make([]int, 0)
	for _, v1 := range data1 {
		for _, v2 := range data2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}
	fmt.Println(data1, data2, result)
}
