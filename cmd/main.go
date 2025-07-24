package main

import (
	"context"
	"fmt"
	"goworkerpool/internal/task"
)

func main() {
	fmt.Println("hello")

	t1 := "http_get:https://go.dev"
	t2 := "compute:factorial:25"

	tp := task.NewTaskParser()
	task1, _ := tp.ParseTask(t1)
	task2, _ := tp.ParseTask(t2)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "helo", "helo")

	fmt.Println(task1)
	task1.Execute(ctx)
	fmt.Println(task2)
	task2.Execute(ctx)
}
