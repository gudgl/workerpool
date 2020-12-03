package main

import (
	"context"
	"fmt"

	"github.com/gudgl/workerpool"
)

func main() {
	client, _ := workerpool.New(
		context.Background(),
		workerpool.WithNumOfWorkers[struct{}](5),
	)

	client.Go()

	for i := 0; i < 100; i++ {
		i := i
		client.PublishJob(func() struct{} {
			if i%2 == 0 {
				fmt.Printf("odd %d\n", i)
				return struct{}{}
			}
			fmt.Printf("even %d\n", i)
			return struct{}{}
		})
	}

	client.Wait()
}
