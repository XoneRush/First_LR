package funcs

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func NewCircle(r int) circle {
	res := circle{
		radius: r,
	}
	return res
}

func (circ circle) Area() {
	fmt.Println("Площадь круга: ", roundFloat(math.Pow(float64(circ.radius), 2)*math.Pi, 2))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
