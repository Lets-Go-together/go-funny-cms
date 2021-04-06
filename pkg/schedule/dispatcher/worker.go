package dispatcher

type Worker interface {
	handle(task *Task)
	AcceptType() []*TaskType
	Initialize()
}
