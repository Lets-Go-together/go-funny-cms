package schedule

type MySQLTaskBroker struct {
	taskChan chan *Task
}

func (that *MySQLTaskBroker) Launch() {
	panic("implement me")
}

func (that *MySQLTaskBroker) RestoreTask() {
	panic("implement me")
}

func (that *MySQLTaskBroker) StartConsuming() (ch <-chan *Task, close chan int) {
	panic("implement me")
}

func (that *MySQLTaskBroker) UpdateTask(info *Task) error {
	panic("implement me")
}

func (that *MySQLTaskBroker) StopTask(id int) error {
	panic("implement me")
}

func (that *MySQLTaskBroker) StartTask(id int) error {
	panic("implement me")
}

func (that *MySQLTaskBroker) AddTask(info *Task) (*Task, error) {
	panic("implement me")
}

func (that *MySQLTaskBroker) DeleteTask(id int) error {
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
