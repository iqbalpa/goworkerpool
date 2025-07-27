package worker

import (
	"context"
	"fmt"
	"goworkerpool/pkg"
)

func StartWorker(id int, queue <-chan pkg.Task, ctx context.Context) {
	fmt.Println("Started worker", id)

	for {
		select {

		case <-ctx.Done():
			return

		case t, ok := <-queue:
			if !ok {
				return
			}
			t.Execute(ctx)
		}
	}
}
