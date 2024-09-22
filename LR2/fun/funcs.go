package fun

import (
	"fmt"
	"os"
)

func init() {}

func IsOdd() {
	number := 0

	fmt.Fscan(os.Stdin, &number)
	if number%2 == 0 {
		fmt.Println("Числое четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
