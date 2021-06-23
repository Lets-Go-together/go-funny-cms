package schedule

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"gocms/pkg/schedule/log"
	"strconv"
	"time"
)

const (
	keyTaskCron     = "SCHEDULE:TASKS"
	taskIdIncrement = "id_increment"
)

// RedisTaskBroker 是基于 Redis 实现的任务管理中间人
type RedisTaskBroker struct {
	redis                   *redis.Client
	taskCh                  chan *Task
	closeCh                 chan int
	redisPollDurationSecond time.Duration
	idTaskMap               map[int]*Task
}

func NewRedisTaskBroker(client *redis.Client) TaskBroker {
	var r TaskBroker = &RedisTaskBroker{
		redis:                   client,
		taskCh:                  make(chan *Task),
		closeCh:                 make(chan int),
		redisPollDurationSecond: 3,
	}
	return r
}

func (that *RedisTaskBroker) Launch() {
	exist, _ := that.redis.HExists(keyTaskCron, taskIdIncrement).Result()
	if !exist {
		// 初始化任务 ID 自增键
		that.redis.HSet(keyTaskCron, taskIdIncrement, 0)
	}
}

func (that *RedisTaskBroker) StartConsuming() (<-chan *Task, chan int) {

	go func() {
		var taskBroker TaskBroker = that
		for {
			taskResult, err := that.redis.HGetAll(keyTaskCron).Result()
			that.removeKeyIncrement(taskResult)
			if err != nil {
				log.Err("broker/StartConsuming", err)
				// Process err
				continue
			}

			for _, item := range taskResult {

				var task Task
				// log.I("broker/StartConsuming", item)

				if err = json.Unmarshal([]byte(item), &task); err != nil {
					// Process err
					log.Err("broker/StartConsuming", err)
					continue
				}

				if task.StateInChange() {
					task.broker = &taskBroker
					task.init()
					select {
					case that.taskCh <- &task:
						log.I("broker/StartConsuming", "new task:"+task.String())
					case <-that.closeCh:
						close(that.taskCh)
						log.D("broker/StartConsuming", "task channel closed")
						return
					}
				}
			}
			time.Sleep(time.Second * that.redisPollDurationSecond)
		}
	}()

	return that.taskCh, that.closeCh
}

func (that *RedisTaskBroker) AddTask(task *Task) (*Task, error) {
	log.D("broker/AddTask", task.String())
	task.ExecuteInfo.CreateNow()
	incr, er := that.redis.HIncrBy(keyTaskCron, taskIdIncrement, 1).Result()
	if er != nil {
		return nil, er
	}
	task.Id = int(incr)
	j, err := json.Marshal(task)
	if err != nil {
		return nil, er
	}
	id := strconv.Itoa(task.Id)
	_, r := that.redis.HSet(keyTaskCron, id, string(j)).Result()
	if r != nil {
		return nil, er
	}
	return task, nil
}

func (that *RedisTaskBroker) DeleteTask(id int) error {
	task := that.QueryTaskById(id)
	if task == nil {
		return errors.New("task does not exist")
	}
	if task.State == TaskStateDeleting {
		_, err := that.redis.HDel(keyTaskCron, strconv.Itoa(id)).Result()
		return err
	}
	return task.ChangeState(TaskStateDeleting)
}

func (that *RedisTaskBroker) UpdateTask(task *Task) error {
	log.D("broker/UpdateTask", "update task: name:", task.Name, ", id:", task.Id, ", state:", task.State)
	j, err := json.Marshal(task)
	if err != nil {
		return err
	}
	id := strconv.Itoa(task.Id)
	// TODO 2021年6月23日10:42:20 notify worker
	_, r := that.redis.HSet(keyTaskCron, id, string(j)).Result()
	if r != nil {
		return r
	}
	return nil
}

func (that *RedisTaskBroker) RestoreTask() {

	idTaskMap, err := that.redis.HGetAll(keyTaskCron).Result()
	if err != nil {
		log.Err("broker/RestoreTask", err)
	}

	for k, item := range idTaskMap {
		if k == taskIdIncrement {
			continue
		}
		var task *Task

		if err = json.Unmarshal([]byte(item), &task); err != nil {
			log.Err("broker/RestoreTask", err)
			continue
		}
		var s TaskBroker = that
		task.broker = &s
		task.init()

		if task.State == TaskStateStopping {
			task.State = TaskStateStopped
			err = that.UpdateTask(task)
			if err != nil {
				panic(err)
			}
		} else if task.State == TaskStateDeleting {
			// delete
			that.redis.HDel(keyTaskCron, strconv.Itoa(task.Id))

		} else if task.State == TaskStateRunning || task.State == TaskStateStarting ||
			task.State == TaskStateRebooting || task.State == TaskStateInitialize {
			// restore
			that.taskCh <- task
		}

		log.D("broker/RestoreTask", "name:", task.Name, ", id:", task.Id, ", state:", task.State)
	}
}

func (that *RedisTaskBroker) StopTask(id int) error {

	j, err := that.redis.HGet(keyTaskCron, strconv.Itoa(id)).Result()
	if err != nil {
		panic(err)
	}

	var task Task
	err = json.Unmarshal([]byte(j), &task)
	if err != nil {
		panic(err)
	}

	task.State = TaskStateStopping
	return that.UpdateTask(&task)
}

func (that *RedisTaskBroker) StartTask(id int) error {

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
	return that.UpdateTask(&task)
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
		log.Err("broker/QueryTaskById", err)
		return nil
	}

	var task Task
	if err = json.Unmarshal([]byte(j), &task); err != nil {
		log.Err("broker/QueryTaskById", err)
		return nil
	}
	var b TaskBroker = that
	task.broker = &b
	task.init()
	return &task
}

func (that *RedisTaskBroker) QueryTaskByName(name string) []*Task {
	// TODO 2021年6月21日16:07:21
	return []*Task{}
}

func (that *RedisTaskBroker) QueryAllTask() []*Task {

	idTaskMap, err := that.redis.HGetAll(keyTaskCron).Result()
	that.removeKeyIncrement(idTaskMap)

	if err != nil {
		log.Err("broker/QueryAllTask", err)
	}

	var allTask []*Task

	for _, item := range idTaskMap {
		var task *Task

		if err = json.Unmarshal([]byte(item), &task); err != nil {
			log.Err("broker/QueryAllTask", err)
			log.Err("broker/QueryAllTask", errors.New(item))
			continue
		}
		var s TaskBroker = that
		task.broker = &s
		task.init()

		allTask = append(allTask, task)
	}

	return allTask
}

func (that *RedisTaskBroker) removeKeyIncrement(redisResult map[string]string) {
	delete(redisResult, taskIdIncrement)
}
