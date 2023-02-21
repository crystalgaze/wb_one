package main

import (
	"fmt"
	"math/big"
)

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	var a, b = big.NewInt(3986490568048564564), big.NewInt(5670982760256092746)
	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
	fmt.Println(Mul(a, b))
	fmt.Println(Div(a, b))
}
