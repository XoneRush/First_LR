package main

import (
	"LR4/functions"
	"fmt"
)

func main() {
	//1
	fmt.Println()
	m := functions.MapInit()

	//2
	fmt.Println()
	fmt.Println("Средний возраст: ", functions.AverAge(m))

	//3
	delete(m, "Andy")
	fmt.Println("После удаления Andy: ", m)

	//4
	functions.ToUppercase()

	//5
	functions.SumOfSeveral()

	//6
	functions.ReverArray()

	fmt.Println()
}
