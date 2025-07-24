package queue

import (
	"context"
	"goworkerpool/pkg"
)

type TaskQueue struct {
	Tasks chan pkg.Task
}

func NewQueue() *TaskQueue {
	return &TaskQueue{
		Tasks: make(chan pkg.Task, 10),
	}
}

func (tq *TaskQueue) AddNewTask(ctx context.Context, task pkg.Task) error {
	tq.Tasks <- task
	return nil
}
