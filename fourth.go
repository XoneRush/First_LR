package main

import (
	"fmt"
	"os"
)

func fourth() {
	var (
		first, second int = 0, 0
		symbol        string
		result        int
	)

	fmt.Print("\nВведите первое число: ")
	fmt.Fscan(os.Stdin, &first)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &second)

	fmt.Print("Введите операцию(+, -, *, /): ")
	fmt.Fscan(os.Stdin, &symbol)

	switch symbol {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		if second == 0 {
			fmt.Println("Деление на 0 невозможно")
			return
		}
		result = first / second
	default:
		result = 0
		fmt.Println("Неверный ввод")
	}
	fmt.Println(result)
}
