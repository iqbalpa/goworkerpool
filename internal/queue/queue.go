package queue

import (
	"context"
	"goworkerpool/pkg"
)

type TaskQueue struct {
	tasks chan pkg.Task
}

func NewQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make(chan pkg.Task),
	}
}

func (tq *TaskQueue) AddNewTask(ctx context.Context) error {
	var task pkg.Task = ctx.Value("task").(pkg.Task)
	tq.tasks <- task
	return nil
}
