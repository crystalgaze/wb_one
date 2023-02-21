package main

import "fmt"

type Human struct {
	Name     string
	LastName string
	Age      int
}

func (human Human) Greet() {
	fmt.Printf("Я %s %s и мне %d лет", human.Name, human.LastName, human.Age)
}

type Action struct {
	Human
}

func main() {
	humanStruct := Human{"Ivan", "Ivanov", 99}
	actionStruct := Action{humanStruct}
	actionStruct.Greet()
}
