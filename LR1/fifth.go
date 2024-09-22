package main

import (
	"fmt"
	"os"
)

func fifth() {
	var (
		first, second float64 = 0, 0
	)

	fmt.Print("\nВведите первое число: ")
	fmt.Fscan(os.Stdin, &first)

	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &second)

	fmt.Println(first + second)
}
