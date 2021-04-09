package schedule

type Worker interface {
	Process(task *Task) error
	Start()
	Stop()
	Initialize(handleFunMap *TaskHandleFuncMap)
}
