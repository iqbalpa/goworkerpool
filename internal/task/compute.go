package task

import (
	"context"
	"fmt"
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
	if c.Type == "factorial" {
		res := 1
		for i := range c.Operand {
			res *= (i + 1)
		}
		fmt.Println(res)
		return nil
	}
	return fmt.Errorf("failed to compute")
}
