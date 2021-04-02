package dispatcher

import (
	"gocms/pkg/schedule/task"
	"gocms/pkg/schedule/worker"
	"time"
)

func loadTaskFrom() []task.Task {
	return []task.Task{}
}

type TaskLoader interface {
	LoadTask() []*task.Task
}

type Dispatcher struct {
	taskLoader []TaskLoader
	workers    []worker.Worker
}

func New() *Dispatcher {
	return &Dispatcher{}
}

func (that *Dispatcher) RegisterTaskType(taskType *task.Type) {

}

func (that *Dispatcher) Start() {

}

func (that *Dispatcher) runWorker() {

}

func (that *Dispatcher) runTaskLoaders() {
	for _, loader := range that.taskLoader {
		go func(l TaskLoader) {
			for {
				var tasks = l.LoadTask()
				for _, t := range tasks {
					that.dispatchTask(t)
				}
				time.Sleep(time.Second)
			}
		}(loader)
	}
}

func (that *Dispatcher) dispatchTask(task *task.Task) {

}
