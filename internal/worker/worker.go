package worker

import (
	"context"
	"fmt"
	"goworkerpool/pkg"
	"time"
)

func StartWorker(id int, queue <-chan pkg.Task, ctx context.Context) {
	fmt.Println("Started worker", id)

	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	select {
	case t := <-queue:
		t.Execute(ctx)
	case <-ctx.Done():
		return
	}
}
