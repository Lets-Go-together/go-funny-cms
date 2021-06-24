package schedule

import "github.com/go-redis/redis"

// Scheduler 为任务队列的入口, 负责维护整个队列的对象
type Scheduler struct {
	// 任务执行分发调度器
	dispatcher *dispatcher
	// 任务存储中间人
	broker TaskBroker
	// 任务名称与对应处理函数映射
	taskHandleFuncMap *TaskHandleFuncMap
}

func New(client *redis.Client) *Scheduler {

	broker := NewRedisTaskBroker(client)
	handleFuncMap := newTaskHandleFuncMap()
	d := newDispatcher(broker, handleFuncMap)
	return &Scheduler{
		dispatcher:        d,
		broker:            broker,
		taskHandleFuncMap: handleFuncMap,
	}
}

func (that *Scheduler) Launch() {
	that.dispatcher.Launch()
	that.broker.RestoreTask()
}

// 注册任务并指定改任务的执行函数 TaskHandleFunc
// pattern: 正则表达式
func (that *Scheduler) RegisterTask(pattern string, handleFunc TaskHandleFunc) {
	if err := that.taskHandleFuncMap.SetHandleFunc(pattern, handleFunc); err != nil {
		panic(err)
	}
}

func (that *Scheduler) DeleteTask(id int) error {
	return that.broker.DeleteTask(id)
}

func (that *Scheduler) AddTask(info *Task) (*Task, error) {
	return that.broker.AddTask(info)
}

func (that *Scheduler) QueryTaskByName(name string) []*Task {
	return that.broker.QueryTaskByName(name)
}

func (that *Scheduler) QueryTaskByState(state TaskState) []*Task {
	return that.broker.QueryTaskByState(state)
}

func (that *Scheduler) StopTask(taskId int) error {
	return that.broker.StopTask(taskId)
}

func (that *Scheduler) StartTask(taskId int) error {
	return that.broker.StartTask(taskId)
}

func (that *Scheduler) QueryTaskById(taskId int) *Task {
	return that.broker.QueryTaskById(taskId)
}

func (that Scheduler) UpdateTask(task *Task) error {
	// TODO 2021年6月24日16:05:48 或许可以在数据库中再加一个更新中的 task 表
	that.dispatcher.UpdateTask(task)
	return that.broker.UpdateTask(task)
}

func (that *Scheduler) QueryAllTask() []*Task {
	return that.broker.QueryAllTask()
}
