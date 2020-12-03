package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gudgl/workerpool"
)

const (
	layout = "2006-01-02"
)

var (
	timeStrings = []string{
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"2020-12-12",
		"2020-12-12",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
		"Jan 2, 2006 at 3:04pm (MST)",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12",
		"2020-12-12",
		"Jan 2, 2006 at 3:04pm (MST)",
		"2020-12-12 13:23:45",
		"2020-12-12 13:23:45",
	}
)

type Resp struct {
	TimeStr string
	Time    time.Time
	Err     error
}

func main() {
	client, _ := workerpool.New(
		context.Background(),
		workerpool.WithCollectors(func(r Resp) {
			if r.Err != nil {
				log.Printf("time_string %s failed, err %s", r.TimeStr, r.Err)
				return
			}
			log.Printf("time_string %s parsed successfully -> %s", r.TimeStr, r.Time)
		}),
	)

	client.Go()

	for _, ts := range timeStrings {
		ts := ts
		client.PublishJob(func() Resp {
			t, err := time.Parse(layout, ts)
			if err != nil {
				return Resp{
					TimeStr: ts,
					Err:     fmt.Errorf("parsing %s failed", ts),
				}
			}

			return Resp{
				TimeStr: ts,
				Time:    t,
			}
		})
	}

	client.Wait()
}
