package dispatcher

type CronWorker struct {
	Worker
}

func (that *CronWorker) handle(task *Task) {

}
func (that *CronWorker) AcceptType() []*TaskType {
	return []*TaskType{TaskTypeDefault}
}

func (that *CronWorker) Initialize() {

}
