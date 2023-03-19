# workerpool

-----
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub release](https://img.shields.io/badge/release-v0.0.1-blue)](https://github.com/gudgl/job-workers/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/gudgl/job-workers.svg)](https://pkg.go.dev/github.com/gudgl/workerpool)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/gudgl/workerpool/blob/main/LICENSE)


Workerpool is a library for asynchronous execution of jobs and handling results from them.
It was made for learning purposes and included generics to touch this topic for the first time. 
Even that this is not the best place to use it. Try it out you can run concurrent jobs and handle their results concurrently.
Open an issue and suggest some improvements so I can learn and maybe this will be usefull to someone.

## Install

```textmate
go get github.com/gudgl/workerpool
```

## Usage

Types
```go
// JobHandler represents something that should be done by the worker
type JobHandler[T any] func() T

// RespHandler represents how to handler workers response to a job
type RespHandler[T any] func(T)

// ConfigOpts represents optional parameters for Config
type ConfigOpts[T any] func(*Config[T])
```

To start create a new `Client`
```go
client, ctx := jw.New(
    // send context 
    ctx,
    // send options if needed
)
```

Options
```go
// WithNumOfWorkers is an option to specify the number of concurrent workers
// the defautl is one
func WithNumOfWorkers[T any](num int) ConfigOpts[T]
// WithNumOfCollectors is an option to specify the number of concurrent collectors
// the defautl is one
func WithNumOfCollectors[T any](num int) ConfigOpts[T]
// WithCollectors is an option to specify the number of concurrent collectors
// the defautl is one
func WithCollectors[T any](collector RespHandler[T]) ConfigOpts[T]
```

Next run the `Go` function from the client to start the workers and the collectors

```go
client.Go()
```

Then publish the jobs

```go
for _, job := range jobs{
    client.PublishJob(job JobHandler)
}
```

or publish them all at once

```go
client.PublishJobs(jobs)
```

Last just wait for workers and collectors to finish

```go
client.Wait()
```

That's all you need to do to get it up running

## Collaboration

If you would like to contribute, head on to the [issues](https://github.com/gudgl/workerpool/issues/new) page for tasks that need help.

## Licence

[MIT](https://github.com/gudgl/workerpool/blob/main/LICENSE)
