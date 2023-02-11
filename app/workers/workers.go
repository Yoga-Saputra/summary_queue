package main

import (
	"log"
	"summary/app/helpers"
	"summary/app/tasks"
	"time"

	"github.com/hibiken/asynq"
)

const redisAddr = "localhost:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeReSummaryDelivery, tasks.HandleRefetchSummaryDeliveryTask)
	mux.HandleFunc(tasks.TypeSummaryDelivery, tasks.HandleRealTimeSummaryDeliveryTask)

	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := srv.Run(mux); err != nil {
			panic(err)
		}
	}()

	Scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr: redisAddr,
		},
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	for _, v := range helpers.GetProviderCode() {

		_, error := Scheduler.Register("@every 60s", tasks.NewRealTimeSummaryDeliveryTask(time.Now(), v))

		if error != nil {
			log.Fatal(error)
		}
	}

	if err := Scheduler.Run(); err != nil {
		panic(err)
	}
}
