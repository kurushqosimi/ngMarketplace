package async

import (
	"log"
	"runtime/debug"
	"sync"
)

// BackgroundRunner structure that holds waitgroup to wait all background operation finish.
type BackgroundRunner struct {
	wg *sync.WaitGroup
}

// NewBackgroundRunner - returns BackgroundRunner to run background operation.
func NewBackgroundRunner(wg *sync.WaitGroup) *BackgroundRunner {
	return &BackgroundRunner{wg: wg}
}

// RunAsync - runs asynchronously any function without arguments.
func (r *BackgroundRunner) RunAsync(fn func()) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()

		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic in background: %v\n%s", rec, debug.Stack())
			}
		}()
		fn()
	}()
}
