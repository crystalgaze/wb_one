package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10}
	group := new(sync.WaitGroup)
	for _, number := range input {
		group.Add(1)
		go func(in int) {
			fmt.Println(in * in)
			defer group.Done()
		}(number)
	}
	group.Wait()
}
