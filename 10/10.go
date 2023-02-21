package main

import "fmt"

func main() {
	input := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	output := make(map[int][]float32)
	for _, v := range input {
		k := int(v) / 10 * 10
		output[k] = append(output[k], v)
	}
	fmt.Println(output)
}
