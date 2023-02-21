package main

import (
	"fmt"
	"sync"
)

func main() {
	oddMap := make(map[int]bool, 0)
	group := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(i int) {
			mutex.Lock()
			if i%2 == 0 {
				oddMap[i] = true
			} else {
				oddMap[i] = false
			}
			mutex.Unlock()
			defer group.Done()
		}(i)
	}
	group.Wait()
	fmt.Println(oddMap)
}
