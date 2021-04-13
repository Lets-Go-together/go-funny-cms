package schedule

type TaskProcessor func([]*Task)

type TaskBroker interface {
	Launch()
	RestoreTask()
	StartConsuming(tasks TaskProcessor)
	UpdateTask(info *Task)
	StopTask(id int)
	AddTask(info *Task) *Task
	QueryAllTask() []*Task
	QueryTaskByState(state TaskState) []*Task
	QueryTaskById(taskId int) *Task
}
