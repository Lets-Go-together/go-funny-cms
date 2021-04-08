package dispatcher

type TaskObserverFunc func([]Task)

type TaskSource interface {
	Initialize()
	SubscribeTaskUpdate(tasks TaskObserverFunc)
	UpdateTask(info *TaskInfo)
	RemoveTask(name string)
	QueryAllTask() []*TaskInfo
	QueryTypeByState(state int) []*TaskInfo
	QueryTaskById(taskId uint64) []*TaskInfo
}

type Dispatcher struct {
	taskSource         TaskSource
	workers            []Worker
	taskExecuteFuncMap TaskManager
}

func New() *Dispatcher {
	d := &Dispatcher{}
	d.SetTaskSource(&RedisTaskSource{})
	return d
}

func (that *Dispatcher) SetTaskSource(source TaskSource) {
	that.taskSource = source
}

func (that *Dispatcher) AddWorker(worker Worker) {
	that.workers = append(that.workers, worker)
}

func (that *Dispatcher) AddTask(task TaskInfo, handleFunc TaskHandleFunc) {
	err := that.taskExecuteFuncMap.SetHandleFunc(task, &handleFunc)
	if err != nil {
		panic(err)
	}
}

func (that *Dispatcher) RemoveTask(taskName string) {
	that.taskSource.RemoveTask(taskName)
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
	that.taskSource.Initialize()
	that.taskSource.SubscribeTaskUpdate(func(tasks []Task) {
		that.onTaskArrive(tasks)
	})
}

func (that *Dispatcher) onTaskArrive(tasks []Task) {
	for _, task := range tasks {
		for _, worker := range that.workers {
			err := worker.NewTask(task)
			if err != nil {
				panic(err)
			}
			// dispatch to other worker.
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
