package main

import (
	"fmt"
	"os"
)

func sixth() {
	var (
		first, second, third int = 0, 0, 0
	)

	fmt.Print("\nВведите первое число: ")
	fmt.Fscan(os.Stdin, &first)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &second)

	fmt.Print("Введите третье число: ")
	fmt.Fscan(os.Stdin, &third)

	fmt.Println((first + second + third) / 3)
}
