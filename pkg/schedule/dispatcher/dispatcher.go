package dispatcher

type TaskObserverFunc func([]*Task)

type TaskSource interface {
	Initialize()
	SubscribeTaskArrive(tasks TaskObserverFunc)
	QueryAllTask() []Task
	QueryTaskById(taskId uint64) []Task
}

type Dispatcher struct {
	taskLoader []TaskSource
	workers    []Worker
}

func New() *Dispatcher {

	return &Dispatcher{}
}

func (that *Dispatcher) AddTaskSource(source TaskSource) {
	that.taskLoader = append(that.taskLoader, source)
}

func (that *Dispatcher) AddWorker(worker Worker) {
	that.workers = append(that.workers, worker)
}

func (that *Dispatcher) RegisterTaskType(taskType *TaskType) {

}

func (that *Dispatcher) Start() {
	that.runTaskLoaders()
	that.runWorkers()
}

func (that *Dispatcher) runWorkers() {
	for _, worker := range that.workers {
		worker.Initialize()
	}
}

func (that *Dispatcher) runTaskLoaders() {
	for _, loader := range that.taskLoader {
		loader.SubscribeTaskArrive(func(tasks []*Task) {
			that.onTaskArrive(tasks)
		})
	}
}

func (that *Dispatcher) onTaskArrive(tasks []*Task) {
	for _, task := range tasks {
		for _, worker := range that.workers {
			worker.handle(task)
			break
		}
	}
}

func (that *Dispatcher) QueryTaskById(taskId uint64) []Task {
	return []Task{}
}

func (that *Dispatcher) QueryAllTask() []Task {
	return []Task{}
}
