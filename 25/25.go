package main

import (
	"fmt"
	"time"
)

func Sleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	fmt.Println("Начало")
	Sleep(5)
	fmt.Println("Конец")
}
