package dispatcher

type Worker interface {
	handle(task Task) bool
	Stop()
	StopNow()
	Initialize()
}
