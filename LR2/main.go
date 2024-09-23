package main

import (
	"LR2/fun"
	"fmt"
	"os"
)

func main() {
	number := 0
	fmt.Println("Введите число")
	fmt.Fscan(os.Stdin, &number)

	//1 задание
	fun.IsOdd(number)
	//2
	fmt.Println(fun.DefineSign(number))
	//3
	fun.PrintFromZeroToTen()
	//4
	fmt.Println("Длина строки: ", fun.StringLen("посчитайте меня"))
	//5
	rect := fun.NewRectangle(10, 10)
	fmt.Println("Площадь прямоугольника: ", fun.Square(rect))
	//6
	fmt.Println("Среднее значение двух чисел: ", fun.Average(10, 5))

}
