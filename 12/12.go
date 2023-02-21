package main

import "fmt"

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	output := make(map[string]struct{})
	for _, v := range input {
		k := v
		output[k] = append(output[k], struct{}{})
	}
	fmt.Println(output)
}
