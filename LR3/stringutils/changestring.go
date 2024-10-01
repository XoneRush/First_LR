package stringutils

import (
	"fmt"
	"unicode/utf8"
)

func Changestring(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func LongestInListOfStrings() {
	names := []string{"bob", "katerine", "tom", "alex"}
	max := 0
	var longestStr string

	fmt.Println()
	for _, name := range names {
		len := utf8.RuneCountInString(name)
		if len > max {
			max = len
			longestStr = name
		}
	}
	fmt.Println("Самая длинная строка: ", longestStr)

}
