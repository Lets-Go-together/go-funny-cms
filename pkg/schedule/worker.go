package schedule

type Worker interface {
	NewTask(task Task) error
	Start()
	Stop()
	StopNow()
	Initialize(handleFunMap *TaskHandleFuncMap)
}
