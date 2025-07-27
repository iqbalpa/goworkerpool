package cli

import (
	"context"
	"fmt"
	"goworkerpool/internal/queue"
	"goworkerpool/internal/task"
	"goworkerpool/internal/worker"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// Services
var (
	taskParser *task.TaskParser
	taskQueue  *queue.TaskQueue
	bufferSize int = 10
	wg         sync.WaitGroup
)

// Flags
var (
	taskStr   string
	workerInt int
)

// CLI Commands
var (
	rootCmd = &cobra.Command{
		Use:   "goworkerpool",
		Short: "Local task worker",
		Long:  "Goworkerpool is a local task worker to run it simultaneously",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goworkerpool -- Task worker v0.1.0")
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Goworkerpool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goworkerpool v0.1.0")
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add new task to the worker",
		Long:  "Adding new task to the worker",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			initServices()
			ctx := context.Background()
			ctx, cancelCtx := context.WithTimeout(ctx, 3*time.Second)
			t, _ := taskParser.ParseTask(taskStr)
			taskQueue.AddNewTask(ctx, t)
			for i := 0; i < workerInt; i++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					worker.StartWorker(id, taskQueue.Tasks, ctx)
				}(i)
			}
			wg.Wait()
			cancelCtx()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	addCmd.PersistentFlags().StringVarP(&taskStr, "task", "t", "", "task in string format")
	addCmd.PersistentFlags().IntVarP(&workerInt, "worker", "w", 1, "the number of worker")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addCmd)
}

func initServices() {
	taskParser = task.NewTaskParser()
	taskQueue = queue.NewQueue(bufferSize)
}
