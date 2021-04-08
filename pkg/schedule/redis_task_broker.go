package schedule

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gocms/pkg/config"
	"strconv"
	"time"
)

const (
	keyTaskCron = "SCHEDULE:TASKS"
)

type RedisTaskBroker struct {
	redis *redis.Client
}

func (that *RedisTaskBroker) Initialize() {
	that.redis = config.Redis
}

func (that *RedisTaskBroker) SubscribeTaskUpdate(taskObserverFunc TaskObserverFunc) {
	go func(taskObserverFunc TaskObserverFunc) {
		for {
			r, err := that.redis.HGetAll(keyTaskCron).Result()
			if err != nil {
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
					var s TaskBroker = that
					task.broker = &s
					tasks = append(tasks, &task)
				}
				fmt.Println("==>> " + task.String())
			}
			if nil != tasks && 0 != len(tasks) {
				taskObserverFunc(tasks)
				fmt.Println("publish " + strconv.Itoa(len(tasks)))
			}
			time.Sleep(time.Second * 10)
		}
	}(taskObserverFunc)
}

func (that *RedisTaskBroker) AddTask(info *TaskInfo) {
	if true {
		return
	}
	j, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	_, r := that.redis.HSet(keyTaskCron, info.Name, string(j)).Result()
	if r != nil {
		panic(r)
	}
}

func (that *RedisTaskBroker) UpdateTask(task *TaskInfo) {
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	_, r := that.redis.HSet(keyTaskCron, task.Name, string(j)).Result()
	if r != nil {
		panic(r)
	}
}

func (that *RedisTaskBroker) RestoreTask() {

}

func (that *RedisTaskBroker) StopTask(id int) {

}

func (that *RedisTaskBroker) QueryTypeByState(state int) []*TaskInfo {
	return []*TaskInfo{}
}

func (that *RedisTaskBroker) QueryTaskById(taskId uint64) []*TaskInfo {
	return []*TaskInfo{}
}

func (that *RedisTaskBroker) QueryAllTask() []*TaskInfo {
	return []*TaskInfo{}
}
