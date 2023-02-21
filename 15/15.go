package main

import "fmt"

// разные символы могут занимать количество байт от 1 до 4
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	u := []rune(v)
	justString = string(u[:100])
	//justString = v[:100]
}
func createHugeString(size int) string {
	var res string
	for i := 0; i < size; i++ {
		res += "£"
	}
	return res
}
func main() {
	someFunc()
	fmt.Println(justString)
}
