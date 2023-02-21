package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	group := new(sync.WaitGroup)
	//используем канал, что бы горутина могла сохранять данные  (а как то еще можно из горутины сохранять?)
	channel := make(chan int, len(numbers))
	sum_sq := 0
	for _, number := range numbers {
		group.Add(1)
		go func(num int) {
			defer group.Done()
			channel <- num * num
		}(number)
	}
	group.Wait()
	close(channel)
	for num := range channel {
		sum_sq += num
	}
	fmt.Println(sum_sq)
}
