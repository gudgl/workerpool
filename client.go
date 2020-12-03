package workerpool

import (
	"context"
	"sync"
)

// Client orchestrates the workers and the collectors for T
type Client[T any] interface {
	// Go starts the workers and the collectors.
	Go()
	// PublishJobs sends all given jobs to the job channel.
	PublishJob(job JobHandler[T])
	// Wait closes the channels and waits for all workers and collectors to finish.
	Wait()
	// Close breaks the flow imediately
	Close()
}

// JobHandler represents something that should be done by the worker
type JobHandler[T any] func() T

// RespHandler represents how to handler workers response to a job
type RespHandler[T any] func(T)

// ConfigOpts represents optional parameters for Config
type ConfigOpts[T any] func(*Config[T])

// Config represents a core for this worker pool
// Default values for numWorkers and numCollectors are 1
type Config[T any] struct {
	ctx    context.Context
	cancel context.CancelFunc

	wgWorkers    sync.WaitGroup
	wgCollectors sync.WaitGroup

	numWorkers    int
	numCollectors int

	collector RespHandler[T]

	jobs  chan JobHandler[T]
	resps chan T

	isCollectorSet bool
}

// New creates new workerpool.Config
func New[T any](
	ctx context.Context,
	opts ...ConfigOpts[T],
) (Client[T], context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	cfg := &Config[T]{
		ctx:           ctx,
		cancel:        cancel,
		numWorkers:    1,
		numCollectors: 1,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.collector != nil {
		cfg.isCollectorSet = true
	}

	cfg.jobs = make(chan JobHandler[T], cfg.numWorkers)
	cfg.resps = make(chan T, cfg.numCollectors)

	return cfg, ctx
}

// WithNumOfWorkers is an option to specify the number of concurrent workers
// the defautl is one
func WithNumOfWorkers[T any](num int) ConfigOpts[T] {
	return func(c *Config[T]) {
		c.numWorkers = num
	}
}

// WithNumOfCollectors is an option to specify the number of concurrent collectors
// the defautl is one
func WithNumOfCollectors[T any](num int) ConfigOpts[T] {
	return func(c *Config[T]) {
		c.numCollectors = num
	}
}

// WithCollectors is an option to specify the number of concurrent collectors
// the defautl is one
func WithCollectors[T any](collector RespHandler[T]) ConfigOpts[T] {
	return func(c *Config[T]) {
		c.collector = collector
	}
}

// Go starts the workers and the collectors.
func (wp *Config[T]) Go() {
	var i int

	for i = 0; i < wp.numWorkers; i++ {
		wp.wgWorkers.Add(1)
		go wp.startWorkers()
	}

	if wp.isCollectorSet {
		for i = 0; i < wp.numCollectors; i++ {
			wp.wgCollectors.Add(1)
			go wp.startCollectors()
		}
	}
}

// Wait closes the channels and waits for all workers and collectors to finish.
func (wp *Config[T]) Wait() {
	close(wp.jobs)
	wp.wgWorkers.Wait()

	if wp.isCollectorSet {
		close(wp.resps)
		wp.wgCollectors.Wait()
	}
}

// PublishJobs sends all given jobs to the job channel.
func (wp *Config[T]) PublishJob(job JobHandler[T]) {
	wp.publishJob(job)
}

// publishJob sends a job to the jobs channel.
func (wp *Config[T]) publishJob(job JobHandler[T]) {
	if job != nil {
		select {
		case <-wp.ctx.Done():
			return
		default:
			wp.jobs <- job
		}
	}
}

// startWorkers assigning the worker to listen to the jobs channel.
func (wp *Config[T]) startWorkers() {
	defer wp.wgWorkers.Done()

	for job := range wp.jobs {
		select {
		case <-wp.ctx.Done():
			return
		default:
			resp := job()
			if wp.isCollectorSet {
				wp.resps <- resp
			}
		}
	}
}

// startCollectors assigning the collectors to listen to the errors channel.
func (wp *Config[T]) startCollectors() {
	defer wp.wgCollectors.Done()

	for resp := range wp.resps {
		wp.collector(resp)
	}
}

// Close ends all workers and breaks the flow
func (wp *Config[T]) Close() {
	if wp.cancel != nil {
		wp.cancel()
	}
}
