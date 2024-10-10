package gorutins

import (
	"fmt"
	"math/rand"
	"time"
)

func GeomProg(size int, mult int) int {
	num := 1
	for i := 0; i < size; i++ {
		num *= mult
		fmt.Printf("Num is: %d\n", num)
		time.Sleep(time.Millisecond * 500)
	}
	return num
}

func ArifProg(size int, q int) int {
	sum := 1
	for i := 0; i < size; i++ {
		sum += q
		fmt.Printf("Sum is: %d\n", sum)
		time.Sleep(time.Millisecond * 500)
	}
	return sum
}

func Fibonacchi(size int) int {
	num := 0
	b := 1
	for i := 2; i < size; i++ {
		num, b = b, num+b
		fmt.Printf("Fib num: %d\n", num)
		time.Sleep(time.Millisecond * 500)
	}
	return num
}

func FibonacchiChannel(size int, channel chan int) {
	num := 0
	b := 1
	for i := 2; i < size; i++ {
		num, b = b, num+b
		fmt.Printf("Fib num: %d\n", num)
		time.Sleep(time.Millisecond * 500)
		channel <- num
	}
	close(channel)
}

func GenerateRandomNum(channel chan int) {
	num := rand.Intn(100-10) + 10
	channel <- num
	fmt.Println(num)
	time.Sleep(time.Second)
}

func OddCheck(num int, msg chan string) {
	if num%2 == 0 {
		msg <- "num isnt odd"
		time.Sleep(time.Second)
	} else {
		msg <- "num is odd"
		time.Sleep(time.Second)
	}
}
