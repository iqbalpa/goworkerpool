package worker

import (
	"context"
	"fmt"
	"goworkerpool/pkg"
	"sync"
)

func StartWorker(id int, queue <-chan pkg.Task, wg *sync.WaitGroup, ctx context.Context) {
	fmt.Printf("Started worker %d", id)
}
