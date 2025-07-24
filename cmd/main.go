package main

import (
	"context"
	"fmt"
	"goworkerpool/internal/queue"
	"goworkerpool/internal/task"
	"goworkerpool/internal/worker"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("hello")

	t1 := "http_get:https://go.dev"
	t2 := "compute:factorial:25"

	tp := task.NewTaskParser()
	task1, _ := tp.ParseTask(t1)
	task2, _ := tp.ParseTask(t2)

	ctx := context.Background()
	ctx, cancelCtx := context.WithTimeout(ctx, 10*time.Second)

	fmt.Println(task1)
	fmt.Println(task2)

	tq := queue.NewQueue(10)
	tq.AddNewTask(ctx, task1)
	tq.AddNewTask(ctx, task2)

	n := 5

	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.StartWorker(i, tq.Tasks, ctx)
		}()
	}

	wg.Wait()
	fmt.Println("Cancelling the context...")
	cancelCtx()
}
