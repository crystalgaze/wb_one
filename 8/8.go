package main

import "fmt"

func main() {
	var value int64 = 128
	// 1000 0000
	fmt.Printf("%08b %d \n", value, value)

	i := 3
	value = value | (1 << i) //операция ИЛИ
	// 1000 1000
	// 132
	fmt.Printf("%08b %d \n", value, value)

	i = 7
	value = value &^ (1 << i) // операция И НЕ
	// 1000 1000
	// 132
	fmt.Printf("%08b %d \n", value, value)
}
