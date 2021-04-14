package schedule

// Scheduler 为任务队列的入口, 负责维护整个队列的对象
type Scheduler struct {
	// 任务执行分发调度器
	dispatcher *dispatcher
	// 任务存储中间人
	broker TaskBroker
	// 任务名称与对应处理函数映射
	taskHandleFuncMap *TaskHandleFuncMap
}

func New() *Scheduler {

	broker := NewRedisTaskBroker()
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
func (that *Scheduler) RegisterTask(taskName string, handleFunc TaskHandleFunc) {
	if err := that.taskHandleFuncMap.SetHandleFunc(taskName, handleFunc); err != nil {
		panic(err)
	}
}

func (that *Scheduler) AddTask(info *Task) *Task {
	return that.broker.AddTask(info)
}

func (that *Scheduler) QueryTaskByName(name string) []*Task {
	return that.broker.QueryTaskByName(name)
}

func (that *Scheduler) StopTask(taskId int) {
	that.broker.StopTask(taskId)
}

func (that *Scheduler) StartTask(taskId int) {
	that.broker.StartTask(taskId)
}

func (that *Scheduler) QueryTaskById(taskId int) *Task {
	return that.broker.QueryTaskById(taskId)
}

func (that *Scheduler) QueryAllTask() []*Task {
	return that.broker.QueryAllTask()
}
