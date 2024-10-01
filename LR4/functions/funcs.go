package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MapInit() map[string]int {
	m := make(map[string]int)

	m["Andy"] = 30
	m["Marie"] = 20
	m["Alex"] = 45

	fmt.Print("Карта после добавления людей выглядит так: ", m)
	return m
}

func AverAge(m map[string]int) float32 {
	size := len(m)
	var sum int

	for _, person := range m {
		sum += person
	}

	return float32(sum) / float32(size)

}

func ToUppercase() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fmt.Println(strings.ToUpper(line))
}

func SumOfSeveral() {
	fmt.Print("Введите числа через пробел: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	nums := stringToInt(line)

	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func ReverArray() {
	fmt.Print("Введите числа через пробел: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	nums := stringToInt(line)

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Print(nums[i], " ")
	}
}

func stringToInt(line string) []int {
	nums_str := strings.Split(line, " ")
	var nums []int
	for _, num_str := range nums_str {

		if num, err := strconv.Atoi(num_str); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}
