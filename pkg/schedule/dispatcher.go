package schedule

import (
	"gocms/pkg/schedule/log"
)

// dispatcher 表示任务调度者, 负责接收从 TaskBroker 发布的任务, 并分发给执行者 Worker
type dispatcher struct {
	broker TaskBroker
	// 任务执行者, 暂时只有一个 cron 实现的执行者
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
		worker.Launch(that.handleFuncMap)
	}
	ch, _ := that.broker.StartConsuming()
	go func() {
		for t := range ch {
			that.onTaskArrive(t)
		}
	}()
	//done <- 1
	that.broker.Launch()
}

func (that *dispatcher) UpdateTask(task *Task) {

}

func (that *dispatcher) onTaskArrive(task *Task) {
	for _, worker := range that.workers {
		log.D("dispatcher/onTaskArrive",
			"dispatch task: ", "name=", task.Name, ", state=", task.State, ", id=", task.Id)
		err := worker.Process(task)
		if err != nil {
			log.E("dispatcher/onTaskArrive", err)
		}
		// dispatch to other worker.
		break
	}
}
