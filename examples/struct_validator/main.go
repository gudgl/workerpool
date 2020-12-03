package main

import (
	"context"
	"errors"
	"log"

	"github.com/gudgl/workerpool"
)

var (
	triangles = []Triangle{
		{3, 4, 5},
		{1, 2, 3},
		{3, 5, 4},
		{5, 4, 3},
		{2, 4, 5},
		{5, 7, 12},
	}

	errInvalidTriangle = errors.New("invalid triangle")
)

type Triangle struct {
	a int
	b int
	c int
}

type Resp struct {
	Err error
	Msg string
}

func main() {
	client, _ := workerpool.New(
		context.Background(),
		workerpool.WithCollectors(func(r Resp) {
			if r.Err != nil {
				log.Printf("Err: %s, Message %s", r.Err, r.Msg)
				return
			}
			log.Printf("OK message %s", r.Msg)
		}),
	)

	client.Go()

	for _, t := range triangles {
		t := t
		client.PublishJob(func() Resp {
			if t.a+t.b <= t.c ||
				t.a+t.c <= t.b ||
				t.b+t.c <= t.a {
				return Resp{
					Err: errInvalidTriangle,
					Msg: "not a triangle",
				}
			}

			if t.a == t.b && t.a == t.c {
				return Resp{
					Msg: "equiteral triangle",
				}
			}

			if t.a != t.b &&
				t.a != t.c &&
				t.b != t.c {
				return Resp{
					Msg: "scalene triangle",
				}
			}

			return Resp{
				Msg: "isoceles triangle",
			}
		})
	}

	client.Wait()
}
