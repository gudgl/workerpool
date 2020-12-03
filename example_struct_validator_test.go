package workerpool_test

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

type TriangleResp struct {
	Err error
	Msg string
}

func Example_struct_validation() {
	client, _ := workerpool.New(
		context.Background(),
		workerpool.WithCollectors(func(r TriangleResp) {
			if r.Err != nil {
				log.Printf("Err: %v, Message %s", r.Err, r.Msg)
				return
			}
			log.Printf("OK message %s", r.Msg)
		}),
	)

	client.Go()

	for _, t := range triangles {
		t := t
		client.PublishJob(func() TriangleResp {
			if t.a+t.b <= t.c ||
				t.a+t.c <= t.b ||
				t.b+t.c <= t.a {
				return TriangleResp{
					Err: errInvalidTriangle,
					Msg: "not a triangle",
				}
			}

			if t.a == t.b && t.a == t.c {
				return TriangleResp{
					Msg: "equiteral triangle",
				}
			}

			if t.a != t.b &&
				t.a != t.c &&
				t.b != t.c {
				return TriangleResp{
					Msg: "scalene triangle",
				}
			}

			return TriangleResp{
				Msg: "isoceles triangle",
			}
		})
	}

	client.Wait()
}
