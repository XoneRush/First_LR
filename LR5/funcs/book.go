package funcs

import "fmt"

type book struct {
	name   string
	author string
}

func Create(bookName string, bookAuthor string) book {
	res := book{
		name:   bookName,
		author: bookAuthor,
	}
	return res
}

func (b book) GetInfo() {
	fmt.Println(b.name, ", ", b.author)
}
