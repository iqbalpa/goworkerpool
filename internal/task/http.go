package task

import (
	"context"
	"fmt"
	"net/http"
)

type HttpGetTask struct {
	Url string
}

func NewHttpGetTask(url string) *HttpGetTask {
	return &HttpGetTask{
		Url: url,
	}
}

func (h *HttpGetTask) Name() string {
	s := fmt.Sprintf("Get Request to %s", h.Url)
	return s
}

func (h *HttpGetTask) Execute(ctx context.Context) error {
	res, err := http.Get(h.Url)
	if err != nil {
		fmt.Printf("Failed to send Get request to %s\n%v", h.Url, err)
		return err
	}

	fmt.Println("Succesfully send Get request to", h.Url)
	fmt.Println(res)
	return nil
}
