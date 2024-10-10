package workers

import (
	"fmt"
	"sync"
	"time"
)

// Task
type Task struct {
	Input string
	Res   chan string
}

// Process task
func (t *Task) Process() {
	result := reverse(t.Input)
	t.Res <- result
	time.Sleep(2 * time.Second)
}

func GetRes(results chan string) {
	for result := range results {
		fmt.Println(result)
	}
}

// WorkerPool
type WorkerPool struct {
	Tasks        []Task
	Concurrency  int
	tasksChannel chan Task
	wg           sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChannel {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.tasksChannel = make(chan Task, len(wp.Tasks))
	results := make(chan string)

	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		task.Res = results
		wp.tasksChannel <- task
	}

	go GetRes(results)

	close(wp.tasksChannel)

	wp.wg.Wait()
	close(results)
}
