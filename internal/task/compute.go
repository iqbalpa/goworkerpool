package task

import (
	"context"
	"fmt"
	"time"
)

type ComputeTask struct {
	Type    string
	Operand int
}

func NewComputeTask(types string, operand int) *ComputeTask {
	return &ComputeTask{
		Type:    types,
		Operand: operand,
	}
}

func (c *ComputeTask) Name() string {
	s := fmt.Sprintf("Compute Task\nOperand: %d\nType: %s", c.Operand, c.Type)
	return s
}

func (c *ComputeTask) Execute(ctx context.Context) error {
	res := 1
	for i := 2; i <= c.Operand; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return ctx.Err()
		case <-time.After(25 * time.Millisecond):
			res *= i
		}
	}
	fmt.Println(res)
	return nil
}
