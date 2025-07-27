# Go Worker Pool

A concurrent task processing system built in Go that implements a worker pool pattern for executing various types of tasks efficiently. This project is a command-line interface (CLI) application.

## Features

- **Concurrent Task Processing**: Utilizes goroutines to process multiple tasks simultaneously
- **Task Queue**: Buffered channel-based task queue for efficient task distribution
- **Multiple Task Types**: Supports HTTP requests and computational tasks
- **Context-based Cancellation**: Proper timeout and cancellation handling
- **Extensible Design**: Easy to add new task types through interface implementation
- **CLI Interface**: Simple and easy to use CLI to run the worker

## Architecture

The project follows a clean architecture pattern with the following components:

```
goworkerpool/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── queue/           # Task queue implementation
│   ├── task/            # Task implementations and parser
│   └── worker/          # Worker pool implementation
├── pkg/
│   └── types.go         # Common interfaces and types
└── go.mod
```

### Core Components

- **Task Interface**: Defines the contract for all executable tasks
- **TaskQueue**: Manages task distribution using buffered channels
- **Worker Pool**: Spawns multiple workers to process tasks concurrently
- **Task Parser**: Parses string-based task definitions into executable tasks

## Supported Task Types

### 1. HTTP GET Tasks
Execute HTTP GET requests to specified URLs.

**Format**: `http_get:https://example.com`

### 2. Compute Tasks
Perform computational operations (currently supports factorial calculation).

**Format**: `compute:factorial:25`

## Installation

1. Clone the repository:
```bash
git clone https://github.com/iqbalpa/goworkerpool.git
cd goworkerpool
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

### Basic Example

```bash
# Build the application
go build -o tmp/main cmd/main.go

# Run the application
./tmp/main add -w 5 -t "http_get:https://go.dev" -t "compute:factorial:25"
```

Or run directly:
```bash
go run cmd/main.go add -w 5 -t "http_get:https://go.dev" -t "compute:factorial:25"
```

### CLI Commands

#### `add`
Add new task to the worker.

**Flags**:
- `-w`, `--worker`: The number of worker (default: 1)
- `-t`, `--task`: Task in string format

#### `version`
Print the version number of Goworkerpool.

## Adding New Task Types

To add a new task type, implement the `Task` interface:

```go
type Task interface {
    Name() string
    Execute(ctx context.Context) error
}
```

### Example: Adding a Sleep Task

1. Create the task struct:
```go
type SleepTask struct {
    Duration time.Duration
}

func (s *SleepTask) Name() string {
    return fmt.Sprintf("Sleep Task (%v)", s.Duration)
}

func (s *SleepTask) Execute(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(s.Duration):
        fmt.Println("Sleep completed")
        return nil
    }
}
```

2. Update the task parser to handle the new format:
```go
case "sleep":
    duration, _ := time.ParseDuration(parts[1])
    return &SleepTask{Duration: duration}, nil
```

## Configuration

### Worker Pool Size
Adjust the number of workers by modifying the `-w` or `--worker` flag.

### Queue Size
Change the queue buffer size by modifying the parameter in `queue.NewQueue(size)`.

### Task Timeout
Modify the context timeout in `main.go` or individual task timeouts in the worker implementation.

## Key Design Patterns

- **Worker Pool Pattern**: Manages a fixed number of workers to process tasks
- **Producer-Consumer Pattern**: Task queue acts as a buffer between task producers and consumers
- **Interface Segregation**: Clean separation of concerns through well-defined interfaces
- **Context Propagation**: Proper context handling for cancellation and timeouts

## Dependencies

- Go 1.24.4 or later
- [github.com/spf13/cobra](https://github.com/spf13/cobra)

## Error Handling

The system includes comprehensive error handling:
- Context cancellation and timeout handling
- HTTP request error handling
- Task parsing error handling
- Graceful shutdown on context cancellation

## Performance Considerations

- Buffered channels prevent blocking when adding tasks
- Worker pool size can be tuned based on workload
- Each task has individual timeout handling
- Memory-efficient design with minimal allocations

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## Future Enhancements

- [ ] Add comprehensive test coverage
- [ ] Implement task prioritization
- [ ] Add metrics and monitoring
- [ ] Support for persistent task queues
- [ ] Configuration file support
- [ ] More task types (file operations, database operations, etc.)
- [ ] Graceful shutdown with task completion