package schedule

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"gocms/pkg/config"
	"gocms/pkg/schedule/log"
	"strconv"
	"time"
)

const (
	keyTaskCron = "SCHEDULE:TASKS"
)

type RedisTaskBroker struct {
	redis            *redis.Client
	taskObserverFunc TaskProsessor
}

func (that *RedisTaskBroker) Launch() {
	that.redis = config.Redis
}

func (that *RedisTaskBroker) StartConsuming(taskObserverFunc TaskProsessor) {
	that.taskObserverFunc = taskObserverFunc

	go func(taskObserverFunc TaskProsessor) {

		for {
			r, err := that.redis.HGetAll(keyTaskCron).Result()
			if err != nil {
				log.Err("broker", err)
				// Process err
				continue
			}

			var tasks []*Task
			for _, item := range r {
				var task Task

				if err = json.Unmarshal([]byte(item), &task); err != nil {
					// Process err
					log.Err("broker", err)
					continue
				}

				if task.StateInChange() {
					var s TaskBroker = that
					task.broker = &s
					tasks = append(tasks, &task)
					log.D("broker", "task update: "+task.String())
				}
			}

			if nil != tasks && 0 != len(tasks) {
				taskObserverFunc(tasks)
				log.D("broker", "task count: "+strconv.Itoa(len(tasks)))
			}
			time.Sleep(time.Second * 3)
		}
	}(taskObserverFunc)
}

func (that *RedisTaskBroker) AddTask(task *Task) {
	log.D("broker", "add task: "+task.String())
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	_, r := that.redis.HSet(keyTaskCron, task.Name, string(j)).Result()
	if r != nil {
		panic(r)
	}
}

func (that *RedisTaskBroker) UpdateTask(task *Task) {
	log.D("broker", "update task: "+task.String())
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

func (that *RedisTaskBroker) QueryTypeByState(state int) []*Task {
	return []*Task{}
}

func (that *RedisTaskBroker) QueryTaskById(taskId uint64) []*Task {
	return []*Task{}
}

func (that *RedisTaskBroker) QueryAllTask() []*Task {
	return []*Task{}
}

func (that RedisTaskBroker) name() {

}
