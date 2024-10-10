package main

import (
	"LR6/workers"
)

func main() {
	// 1 задание создание и запуск горутин
	// t := time.Now()

	// go gorutins.GeomProg(5, 2)
	// go gorutins.ArifProg(5, 2)
	// gorutins.Fibonacchi(5)
	// fmt.Printf("Total time: %.2f seconds", time.Since(t).Seconds())

	// 2 задание использование каналов для передачи данных
	// channel := make(chan int)
	// go gorutins.FibonacchiChannel(10, channel)

	// for num := range channel {
	// 	fmt.Println(num)
	// }

	//3 задание применение select для управления каналами

	// intChannel := make(chan int)
	// stringChannel := make(chan string)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		gorutins.GenerateRandomNum(intChannel)
	// 	}
	// 	close(intChannel)
	// }()

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case val, ok := <-intChannel:
	// 		if ok {
	// 			go gorutins.OddCheck(val, stringChannel)
	// 		}
	// 	case msg := <-stringChannel:
	// 		fmt.Println(msg)
	// 	}
	// }

	// time.Sleep(time.Second * 2)

	// 4 задание

	// count := 0
	// var mu sync.Mutex

	// go func() {
	// 	mu.Lock()
	// 	for i := 0; i < 5; i++ {
	// 		count++
	// 		fmt.Println("Первая горутина увеличила счетчик на: ", count)
	// 	}
	// 	mu.Unlock()
	// }()

	// go func() {
	// 	mu.Lock()
	// 	for i := 0; i < 5; i++ {
	// 		count++
	// 		fmt.Println("Вторая горутина увеличила счетчик на: ", count)
	// 	}
	// 	mu.Unlock()
	// }()

	// time.Sleep(time.Second * 2)

	//5 задание
	// requests := make(chan calculator.Request)
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go calculator.Calc(requests, &wg)

	// for i := 0; i < 5; i++ {
	// 	req := calculator.Request{
	// 		First:  float32(i) + 1,
	// 		Second: float32(i) + 2,
	// 		Symbol: "+",
	// 		Res:    make(chan float32),
	// 	}
	// 	requests <- req

	// 	result := <-req.Res
	// 	fmt.Println(result)
	// }
	// close(requests)
	// wg.Wait()

	//6 задание

	tasks := []workers.Task{
		workers.Task{Input: "строка"},
		workers.Task{Input: "еще одна"},
		workers.Task{Input: "другая"},
		workers.Task{Input: "четвертая"},
		workers.Task{Input: "пятая"},
		workers.Task{Input: "шестая"},
	}

	wp := workers.WorkerPool{
		Tasks:       tasks,
		Concurrency: 3,
	}

	wp.Run()

}
