package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	s = strings.ToLower(s)
	dict := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := dict[v]; ok {
			//в ok запишется true если по ключу v есть значение
			return false
		}
		dict[v] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println("abCdefAaf")
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println("abcdefghjkl")
	fmt.Println(isUnique("abcdefghjkl"))
}
