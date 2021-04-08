package schedule

type TaskBroker interface {
	Initialize()
	RestoreTask()
	SubscribeTaskUpdate(tasks TaskObserverFunc)
	UpdateTask(info *TaskInfo)
	StopTask(id int)
	AddTask(info *TaskInfo)
	QueryAllTask() []*TaskInfo
	QueryTypeByState(state int) []*TaskInfo
	QueryTaskById(taskId uint64) []*TaskInfo
}

type TaskCreator interface {
}
