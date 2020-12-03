package workerpool_test

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

type TimeResp struct {
	Str  string
	Time time.Time
	Err  error
}

func main() {
	client, _ := workerpool.New(
		context.Background(),
		workerpool.WithCollectors(func(r TimeResp) {
			if r.Err != nil {
				log.Printf("time_string %s failed, err %s", r.Str, r.Err)
				return
			}
			log.Printf("time_string %s parsed successfully -> %s", r.Str, r.Time)
		}),
	)

	client.Go()

	for _, ts := range timeStrings {
		ts := ts
		client.PublishJob(func() TimeResp {
			t, err := time.Parse(layout, ts)
			if err != nil {
				return TimeResp{
					Str: ts,
					Err: fmt.Errorf("parsing %s failed", ts),
				}
			}

			return TimeResp{
				Str:  ts,
				Time: t,
			}
		})
	}

	client.Wait()
}
