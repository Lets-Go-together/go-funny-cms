package schedule

type TaskProsessor func([]*Task)

type TaskBroker interface {
	Launch()
	RestoreTask()
	StartConsuming(tasks TaskProsessor)
	UpdateTask(info *Task)
	StopTask(id int)
	AddTask(info *Task)
	QueryAllTask() []*Task
	QueryTypeByState(state int) []*Task
	QueryTaskById(taskId uint64) []*Task
}
