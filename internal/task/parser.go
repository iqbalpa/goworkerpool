package task

import (
	"fmt"
	"goworkerpool/pkg"
	"strconv"
	"strings"
)

type TaskParser struct{}

func NewTaskParser() *TaskParser {
	return &TaskParser{}
}

func (tp *TaskParser) ParseTask(raw string) (pkg.Task, error) {
	parts := strings.Split(raw, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("task format is invalid")
	}

	switch parts[0] {

	case "http_get":
		url := fmt.Sprintf("%s:%s", parts[1], parts[2])
		t := NewHttpGetTask(url)
		return t, nil

	case "compute":
		types := parts[1]
		operand, _ := strconv.Atoi(parts[2])
		t := NewComputeTask(types, operand)
		return t, nil

	default:
		return nil, fmt.Errorf("task is invalid")
	}
}
