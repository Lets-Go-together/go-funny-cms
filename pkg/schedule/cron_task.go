package schedule

import "errors"

type CronTask struct {
	*TaskInfo
}

func (that *CronTask) Execute(context *Context) *TaskResult {
	return nil
}

func (that *CronTask) GetInfo() *TaskInfo {
	return that.TaskInfo
}

func (that *CronTask) Context() *Context {
	return nil
}

func (that *CronTask) Log(tag string, log string) {

}

func (that *CronTask) ChangeState(state TaskState) error {
	if that.broker != nil {
		that.State = state
		(*that.broker).UpdateTask(that.TaskInfo)
	} else {
		return errors.New("the task broker is not specify")
	}
	return nil
}

func (that *CronTask) requireSource(s func(source *TaskBroker)) {
	s(that.broker)
}
