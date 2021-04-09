package schedule

type CronTask struct {
	*Task
}

func (that *CronTask) Execute(context *Context) *TaskResult {
	return nil
}

func (that *CronTask) Context() *Context {
	return nil
}

func (that *CronTask) Log(tag string, log string) {

}

func (that *CronTask) requireSource(s func(source *TaskBroker)) {
	s(that.broker)
}
