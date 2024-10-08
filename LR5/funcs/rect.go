package funcs

import "fmt"

type Rectange struct {
	Width  float32
	Height float32
}

func NewRectangle(width float32, height float32) Rectange {
	r := Rectange{Width: width, Height: height}

	return r
}

func (r Rectange) Area() {
	fmt.Println("Площадь прямоугольника: ", r.Width*r.Height)
}
