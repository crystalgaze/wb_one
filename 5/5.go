package main

import (
	"fmt"
	"time"
)

func main() {
	n := 4
	stop := time.After(time.Duration(n) * time.Second)
	//запишет в канал stop Time после истечения n секунд
	c := make(chan int)
	//небуферизованный канал
	go func() {
		for i := 0; i < 10000000; i++ {
			fmt.Printf("Запись %d\n", i)
			c <- i
			time.Sleep(10 * time.Millisecond)
			//без задержки вывод какой то странный
		}
	}()

	go func() {
		for v := range c {
			fmt.Printf("Чтение %d\n", v)
		}
	}()
	<-stop
}
