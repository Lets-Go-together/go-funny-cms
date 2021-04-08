package dispatcher

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"gocms/pkg/config"
	"time"
)

const (
	keyTaskCron = "SCHEDULE:TASKS"
)

type RedisTaskSource struct {
	redis *redis.Client
}

func (that *RedisTaskSource) Initialize() {
	that.redis = config.Redis
}

func (that *RedisTaskSource) SubscribeTaskUpdate(taskObserverFunc TaskObserverFunc) {
	go func(taskObserverFunc TaskObserverFunc) {
		for {
			r, err := that.redis.HGetAll(keyTaskCron).Result()
			if err == nil {
				// NewTask err
				continue
			}
			var tasks []Task
			for _, item := range r {
				var task CronTask
				if err := json.Unmarshal([]byte(item), &task); err != nil {
					// NewTask err
					continue
				}
				if task.StateInChange() {
					tasks = append(tasks, &task)
				}
			}
			taskObserverFunc(tasks)
			time.Sleep(time.Second * 3)
		}
	}(taskObserverFunc)
}

func (that *RedisTaskSource) UpdateTask(task *TaskInfo) {

}
func (that *RedisTaskSource) RemoveTask(taskName string) {

}

func (that *RedisTaskSource) QueryTypeByState(state int) []*TaskInfo {
	return []*TaskInfo{}
}

func (that *RedisTaskSource) QueryTaskById(taskId uint64) []*TaskInfo {
	return []*TaskInfo{}
}

func (that *RedisTaskSource) QueryAllTask() []*TaskInfo {
	return []*TaskInfo{}
}
