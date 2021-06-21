package schedule

type MySQLTaskBroker struct {
	taskProcessor TaskProcessor
}

func (that *MySQLTaskBroker) Launch() {
	panic("implement me")
}

func (that *MySQLTaskBroker) RestoreTask() {
	panic("implement me")
}

func (that *MySQLTaskBroker) StartConsuming(taskProcessor TaskProcessor) {
	that.taskProcessor = taskProcessor
	go func(processor TaskProcessor) {
		// read task from mysql

		// notify to task processor
	}(that.taskProcessor)
}

func (that *MySQLTaskBroker) UpdateTask(info *Task) {
	panic("implement me")
}

func (that *MySQLTaskBroker) StopTask(id int) {
	panic("implement me")
}

func (that *MySQLTaskBroker) StartTask(id int) {
	panic("implement me")
}

func (that *MySQLTaskBroker) AddTask(info *Task) (*Task, error) {
	panic("implement me")
}

func (that *MySQLTaskBroker) QueryAllTask() []*Task {
	panic("implement me")
}

func (that *MySQLTaskBroker) QueryTaskByState(state TaskState) []*Task {
	panic("implement me")
}

func (that *MySQLTaskBroker) QueryTaskByName(name string) []*Task {
	panic("implement me")
}

func (that *MySQLTaskBroker) QueryTaskById(taskId int) *Task {
	panic("implement me")
}
