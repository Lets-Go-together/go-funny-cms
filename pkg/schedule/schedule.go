package schedule

type Scheduler struct {
	dispatcher        *dispatcher
	broker            TaskBroker
	taskHandleFuncMap *TaskHandleFuncMap
}

func New() *Scheduler {

	broker := &RedisTaskBroker{}
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
}

func (that *Scheduler) RegisterTask(taskName string, handleFunc TaskHandleFunc) {
	if err := that.dispatcher.handleFuncMap.SetHandleFunc(taskName, handleFunc); err != nil {
		panic(err)
	}
}

func (that *Scheduler) AddTask(info *Task) {
	that.broker.AddTask(info)
}

func (that *Scheduler) StopTask(taskId int) {
	that.broker.StopTask(taskId)
}

func (that *Scheduler) StartTask(taskId int) {

}

func (that *Scheduler) QueryTaskById(taskId uint64) []Task {
	return []Task{}
}

func (that *Scheduler) QueryAllTask() []Task {
	return []Task{}
}
