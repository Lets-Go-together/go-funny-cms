package dispatcher

type CronTask struct {
	Duration int32
	*TaskEntity
}

func (that *CronTask) Execute() *TaskResult {
	return that.ExecuteFunc()
}

func (that *CronTask) Entity() *TaskEntity {
	return that.TaskEntity
}
