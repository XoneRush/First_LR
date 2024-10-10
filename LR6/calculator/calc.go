package calculator

import (
	"fmt"
	"sync"
)

type Request struct {
	First  float32
	Second float32
	Symbol string
	Res    chan float32
}

func Calc(requests <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()

	for req := range requests {
		var result float32
		switch req.Symbol {
		case "+":
			result = req.First + req.Second
		case "-":
			result = req.First - req.Second
		case "/":
			if req.Second != 0 {
				result = req.First / req.Second
			} else {
				fmt.Println("ошибка, второе число не может быть 0")
				result = 0
			}
		case "*":
			result = req.First * req.Second

		default:
			fmt.Println("ошибка, неверная операция")
			result = 0
		}

		req.Res <- result
	}
}
