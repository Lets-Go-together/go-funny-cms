package config

import (
	"gocms/pkg/schedule"
	"gocms/pkg/schedule/log"
)

func InitScheduler() {
	Scheduler = schedule.New(Redis)

	Scheduler.RegisterTask(".", func(ctx *schedule.Context) error {
		log.D("TaskRun", ctx.Task)
		return nil
	})
	Scheduler.Launch()
}
