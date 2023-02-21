package main

import (
	"fmt"
	"reflect"
)

func getType(in interface{}) string {
	return reflect.TypeOf(in).String()
}
func main() {
	var a, b, c, d, e = 456, float32(32.456), float64(-8.192), true, "foo"
	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
	fmt.Println(getType(e))

}
