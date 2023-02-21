package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	input := []int{1, 2, 3, 4, 5, 6}
	group := new(sync.WaitGroup)

	group.Add(1)
	go func() {
		for _, number := range input {
			ch1 <- number
		}
		defer close(ch1)
		defer group.Done()
	}()

	group.Add(1)
	go func() {
		for number := range ch1 {
			ch2 <- number * 2
		}
		defer close(ch2)
		defer group.Done()
	}()

	group.Add(1)
	go func() {
		for res := range ch2 {
			fmt.Println(res)
		}
		defer group.Done()
	}()

	group.Wait()
}
