package schedule

type TaskObserverFunc func([]Task)

type dispatcher struct {
	broker        TaskBroker
	workers       []Worker
	handleFuncMap *TaskHandleFuncMap
}

func newDispatcher(broker TaskBroker, funcMap *TaskHandleFuncMap) *dispatcher {
	return &dispatcher{
		broker:        broker,
		workers:       []Worker{&CronWorker{}},
		handleFuncMap: funcMap,
	}
}

func (that *dispatcher) run() {
	that.runTaskLoaders()
	that.runWorkers()
}

func (that *dispatcher) runWorkers() {
	for _, worker := range that.workers {
		worker.Initialize(that.handleFuncMap)
	}
}

func (that *dispatcher) runTaskLoaders() {
	that.broker.Initialize()
	that.broker.SubscribeTaskUpdate(func(tasks []Task) {
		that.onTaskArrive(tasks)
	})
}

func (that *dispatcher) onTaskArrive(tasks []Task) {
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
