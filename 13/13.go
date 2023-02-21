package main

import "fmt"

func main() {
	a := 100
	b := 500
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

}
