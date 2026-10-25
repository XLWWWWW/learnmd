// Package channellab contains the exercises described in docs/03-exercises.md.
package channellab

import "errors"

var (
	ErrInvalidWorkers   = errors.New("workers must be positive")
	ErrInvalidBatchSize = errors.New("batch size must be positive")
	ErrNoJobs           = errors.New("no jobs")
)

// Result ties a result (including a per-job error) to its original input.
type Result struct {
	Input int
	Value int
	Err   error
}
