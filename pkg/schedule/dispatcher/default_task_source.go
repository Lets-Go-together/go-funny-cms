package dispatcher

import (
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"gocms/pkg/config"
	"time"
)

type CronTaskSource struct {
	TaskSource
	redis *redis.Client
	cron  *cron.Cron
}

func (that *CronTaskSource) Initialize() {
	that.redis = config.Redis
}

func (that *CronTaskSource) SubscribeTaskArrive(taskObserverFunc TaskObserverFunc) {
	go func(taskObserverFunc TaskObserverFunc) {
		for {
			time.Sleep(time.Second * 3)
		}
	}(taskObserverFunc)
}

func (that *CronTaskSource) QueryTaskById(taskId uint64) []Task {
	return []Task{}
}

func (that *CronTaskSource) QueryAllTask() []Task {
	return []Task{}
}
