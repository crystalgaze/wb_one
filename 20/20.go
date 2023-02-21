package main

import (
	"fmt"
	"strings"
)

func reverseWordsInString(str string) string {
	var in = strings.Split(str, " ")
	//for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
	//	in[i], in[j] = in[j], in[i]
	//}
	//return strings.Join(in, " ")
	var out string
	for i := len(in) - 1; i >= 0; i-- {
		out += in[i] + " "
	}
	return out
}

func main() {
	fmt.Println("snow dog sun")
	fmt.Println(reverseWordsInString("snow dog sun"))
}
