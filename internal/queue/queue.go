package queue

import (
	"context"
	"fmt"
	"goworkerpool/pkg"
)

type TaskQueue struct {
	Tasks chan pkg.Task
}

func NewQueue(queueSize int) *TaskQueue {
	return &TaskQueue{
		Tasks: make(chan pkg.Task, queueSize),
	}
}

func (tq *TaskQueue) AddNewTask(ctx context.Context, task pkg.Task) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("timed out")
	case tq.Tasks <- task:
		fmt.Println("Added new task to the queue")
		return nil
	}
}
