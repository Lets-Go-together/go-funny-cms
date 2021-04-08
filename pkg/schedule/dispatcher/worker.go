package dispatcher

type Worker interface {
	NewTask(task Task) error
	Start()
	Stop()
	StopNow()
	Initialize()
}
