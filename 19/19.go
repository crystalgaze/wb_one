package main

import "fmt"

func reverseString(str string) string {
	var in = []rune(str)
	var out []rune
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return string(out)
}

func main() {
	fmt.Println("главрыба")
	fmt.Println(reverseString("главрыба"))
	fmt.Println("◀◁◂◃◄◅◆◇◈◉◊○◌◍◎●")
	fmt.Println(reverseString("◀◁◂◃◄◅◆◇◈◉◊○◌◍◎●"))
}
