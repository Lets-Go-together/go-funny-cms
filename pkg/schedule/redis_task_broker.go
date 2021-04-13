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
	keyTaskCron     = "SCHEDULE:TASKS"
	fieldTaskIdIncr = "id_increment"
)

type RedisTaskBroker struct {
	redis         *redis.Client
	taskProcessor TaskProcessor
}

func NewRedisTaskBroker() TaskBroker {
	var r TaskBroker = &RedisTaskBroker{
		redis:         config.Redis,
		taskProcessor: nil,
	}
	return r
}

func (that *RedisTaskBroker) Launch() {
	exist, _ := that.redis.HExists(keyTaskCron, fieldTaskIdIncr).Result()
	if !exist {
		that.redis.HSet(keyTaskCron, fieldTaskIdIncr, 0)
	}
}

func (that *RedisTaskBroker) StartConsuming(taskObserverFunc TaskProcessor) {
	that.taskProcessor = taskObserverFunc

	go func(taskObserverFunc TaskProcessor) {

		for {
			r, err := that.redis.HGetAll(keyTaskCron).Result()
			if err != nil {
				log.Err("broker/StartConsuming", err)
				// Process err
				continue
			}

			var tasks []*Task
			for k, item := range r {
				if k == fieldTaskIdIncr {
					continue
				}

				var task Task
				log.D("broker/StartConsuming", item)

				if err = json.Unmarshal([]byte(item), &task); err != nil {
					// Process err
					log.Err("broker/StartConsuming", err)
					continue
				}

				if task.StateInChange() {
					var s TaskBroker = that
					task.broker = &s
					task.init()
					tasks = append(tasks, &task)
					log.D("broker/StartConsuming", "task consume: "+task.String())
				}
			}

			if nil != tasks && 0 != len(tasks) {
				taskObserverFunc(tasks)
				log.D("broker/StartConsuming", "task count: "+strconv.Itoa(len(tasks)))
			}

			time.Sleep(time.Second * 3)
		}
	}(taskObserverFunc)
}

func (that *RedisTaskBroker) AddTask(task *Task) *Task {
	log.D("broker/AddTask", "add task: "+task.String())
	task.executeInfo.CreateNow()
	incr, er := that.redis.HIncrBy(keyTaskCron, fieldTaskIdIncr, 1).Result()
	if er != nil {
		panic(er)
	}
	task.Id = int(incr)
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	id := strconv.Itoa(task.Id)
	_, r := that.redis.HSet(keyTaskCron, id, string(j)).Result()
	if r != nil {
		panic(r)
	}
	return task
}

func (that *RedisTaskBroker) UpdateTask(task *Task) {
	log.D("broker/UpdateTask", "update task: "+task.String())
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	id := strconv.Itoa(task.Id)
	_, r := that.redis.HSet(keyTaskCron, id, string(j)).Result()
	if r != nil {
		panic(r)
	}
}

func (that *RedisTaskBroker) RestoreTask() {

	idTaskMap, err := that.redis.HGetAll(keyTaskCron).Result()
	if err != nil {
		log.Err("broker/RestoreTask", err)
	}

	for k, item := range idTaskMap {
		if k == fieldTaskIdIncr {
			continue
		}
		var task *Task

		if err = json.Unmarshal([]byte(item), &task); err != nil {
			log.Err("broker/RestoreTask", err)
			continue
		}
		var s TaskBroker = that
		task.broker = &s
		task.executeInfo = &ExecuteInfo{}

		if task.State == TaskSateStopping || task.State == TaskStateStopped {
			task.State = TaskStateStopped
			that.UpdateTask(task)

		} else if task.State == TaskStateDeleting {
			// delete
			that.redis.HDel(keyTaskCron, strconv.Itoa(task.Id))

		} else if task.State == TaskStateRunning || task.State == TaskStateStarting ||
			task.State == TaskStateRebooting || task.State == TaskStateInitialize {
			// run
			that.taskProcessor([]*Task{task})
		}

		log.D("broker/RestoreTask", "task restore: name=%taskState, current_state=%d", task.Name, task.State)
	}
}

func (that *RedisTaskBroker) StopTask(id int) {

	j, err := that.redis.HGet(keyTaskCron, strconv.Itoa(id)).Result()
	if err != nil {
		panic(err)
	}

	var task Task
	err = json.Unmarshal([]byte(j), &task)
	if err != nil {
		panic(err)
	}

	task.State = TaskSateStopping
	that.UpdateTask(&task)
}

func (that *RedisTaskBroker) StartTask(id int) {

	j, err := that.redis.HGet(keyTaskCron, strconv.Itoa(id)).Result()
	if err != nil {
		panic(err)
	}

	var task Task
	err = json.Unmarshal([]byte(j), &task)
	if err != nil {
		panic(err)
	}
	task.State = TaskStateStarting
	that.UpdateTask(&task)
}

func (that *RedisTaskBroker) QueryTaskByState(state TaskState) []*Task {

	var res []*Task
	for _, t := range that.QueryAllTask() {
		if t.State == state {
			res = append(res, t)
		}
	}
	return res
}

func (that *RedisTaskBroker) QueryTaskById(taskId int) *Task {

	t := strconv.Itoa(taskId)
	j, err := that.redis.HGet(keyTaskCron, t).Result()
	if err != nil {
		return nil
	}

	var task Task
	if err = json.Unmarshal([]byte(j), &task); err != nil {
		log.Err("broker", err)
		return nil
	}
	return &task
}

func (that *RedisTaskBroker) QueryAllTask() []*Task {

	idTaskMap, err := that.redis.HGetAll(keyTaskCron).Result()
	if err != nil {
		log.Err("broker", err)
	}

	var allTask []*Task

	for _, item := range idTaskMap {
		var task *Task

		if err = json.Unmarshal([]byte(item), &task); err != nil {
			log.Err("broker", err)
			continue
		}
		var s TaskBroker = that
		task.broker = &s

		allTask = append(allTask, task)
	}

	return allTask
}
