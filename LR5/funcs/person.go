package funcs

import "fmt"

type person struct {
	personName string
	personAge  int
}

func NewPerson(name string, age int) person {
	result := person{
		personName: name,
		personAge:  age,
	}
	return result
}

func GetInfo(p person) {
	fmt.Println(p.personName, ", ", p.personAge)
}

func Birthday(p *person) {
	(*p).personAge++
}
