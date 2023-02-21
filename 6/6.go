package main

import "fmt"

// 1 способ горутина звершит сама себя
// func main() {
// 	stop := make(chan int)
// 	go func() {
// 		select {
// 		case stop <- 1:
// 			return
// 		}
// 	}()
// 	<-stop
// }

// 2 способ горутину завершит сигнал из вне
func main() {
	ch := make(chan int)
	go func() {
		counter := 0
		for {
			select {
			case ch <- counter:
				counter++
				//полезная нагрузка
			case <-ch:
				return
				//если удалось прочитать из канала соответственно в него что то записали, то выходим
			}
		}
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 1
}
