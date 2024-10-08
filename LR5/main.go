package main

import (
	"LR5/funcs"
)

func main() {
	//1 и 2 задание
	john := funcs.NewPerson("John", 20)
	funcs.Birthday(&john)
	funcs.GetInfo(john)

	//3
	circ := funcs.NewCircle(5)
	circ.Area()

	//4
	rect := funcs.NewRectangle(10, 20)
	rect.Area()

	//5
	shapes := []funcs.Shape{
		circ, rect,
	}

	funcs.AreaForSlice(shapes)

	//6
	book := funcs.Create("Красный смех", "Андреев")
	book.GetInfo()
}
