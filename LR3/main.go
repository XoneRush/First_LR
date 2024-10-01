package main

import (
	"LR3/mathutils"
	"LR3/stringutils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1 задание
	fmt.Println(mathutils.Factorial(5))

	//2 задание
	num := 0
	fmt.Println("Введите число")
	fmt.Fscan(os.Stdin, &num)
	fmt.Println("Факториал введенного числа: ", mathutils.Factorial(num))

	//3 задание
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(stringutils.Changestring(line))

	//4 задание
	var arr [5]int
	arr = mathutils.ArrOfInts()
	for i := range arr {
		fmt.Println(arr[i])
	}
	//5 задание
	mathutils.List()

	//6 задание
	stringutils.LongestInListOfStrings()
}
