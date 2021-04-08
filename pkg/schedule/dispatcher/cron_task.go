package dispatcher

type CronTask struct {
	Duration int32
	*TaskInfo
}

func (that *CronTask) Execute() *TaskResult {
	return nil
}

func (that *CronTask) Entity() *TaskInfo {
	return that.TaskInfo
}
