package schedule

import (
	"gocms/pkg/schedule/log"
)

type dispatcher struct {
	broker        TaskBroker
	workers       []Worker
	handleFuncMap *TaskHandleFuncMap
}

func newDispatcher(
	broker TaskBroker,
	funcMap *TaskHandleFuncMap,
) *dispatcher {
	return &dispatcher{
		broker:        broker,
		workers:       []Worker{&CronWorker{}},
		handleFuncMap: funcMap,
	}
}

func (that *dispatcher) Launch() {
	for _, worker := range that.workers {
		worker.Initialize(that.handleFuncMap)
	}
	that.broker.StartConsuming(func(tasks []*Task) {
		that.onTaskArrive(tasks)
	})
	that.broker.Launch()
}

func (that *dispatcher) onTaskArrive(tasks []*Task) {
	for _, task := range tasks {
		for _, worker := range that.workers {
			log.D("dispatcher/onTaskArrive",
				"dispatch task: name=%s, state=%s, id=%s", task.Name, task.StateName(), task.Id)
			err := worker.Process(task)
			if err != nil {
				panic(err)
			}
			// dispatch to other worker.
			break
		}
	}
}
