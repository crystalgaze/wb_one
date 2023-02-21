package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

var numOfWorkers int = 4

func worker(id int, ch <-chan interface{}, group *sync.WaitGroup) {
	defer group.Done()
	for job := range ch {
		fmt.Printf("worker %d start %v\n", id, job)
		fmt.Printf("worker %d finish %v\n", id, job)
	}
}

func main() {
	ch := make(chan interface{}, 10)
	group := new(sync.WaitGroup)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()
	for i := 1; i <= numOfWorkers; i++ {
		group.Add(1)
		go worker(i, ch, group)
	}

	for {
		select {
		case <-ctx.Done():
			close(ch)
			group.Wait()
			//воркеры безопасно доделают работу
			return
		default:
			ch <- []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
			ch <- "строкаСтрока"
			ch <- 7.75437
			ch <- 24562347
			ch <- true
		}
	}
}
