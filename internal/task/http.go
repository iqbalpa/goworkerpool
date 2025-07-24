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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.Url, nil)
	if err != nil {
		return err
	}

	resCh := make(chan *http.Response, 1)
	errCh := make(chan error, 1)
	go func() {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		resCh <- res
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
		return ctx.Err()
	case err := <-errCh:
		return err
	case res := <-resCh:
		defer res.Body.Close()
		fmt.Println(res)
		return nil
	}
}
